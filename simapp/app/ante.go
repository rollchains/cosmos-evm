package app

import (
	"errors"

	"math"

	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	"github.com/cosmos/ibc-go/v8/modules/core/keeper"

	circuitante "cosmossdk.io/x/circuit/ante"
	circuitkeeper "cosmossdk.io/x/circuit/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	evmante "github.com/sei-protocol/sei-chain/x/evm/ante"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper       *keeper.Keeper
	CircuitKeeper   *circuitkeeper.Keeper
	EVMKeeper       *evmkeeper.Keeper
	LatestCtxGetter func() sdk.Context

	BypassMinFeeMsgTypes []string
}

const EVMAssociatePriority = math.MaxInt64 - 100

// NewAnteHandler constructor
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errors.New("account keeper is required for ante builder")
	}
	if options.BankKeeper == nil {
		return nil, errors.New("bank keeper is required for ante builder")
	}
	if options.SignModeHandler == nil {
		return nil, errors.New("sign mode handler is required for ante builder")
	}
	if options.CircuitKeeper == nil {
		return nil, errors.New("circuit keeper is required for ante builder")
	}
	if options.EVMKeeper == nil {
		return nil, errors.New("evm keeper is required for ante builder")
	}
	if options.LatestCtxGetter == nil {
		return nil, errors.New("latest context getter is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		// TODO: evmtypes.MsgAssociate should be glassless
		// antedecorators.NewGaslessDecorator([]sdk.AnteFullDecorator{ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.ParamsKeeper.(paramskeeper.Keeper), options.TxFeeChecker)}, *options.OracleKeeper, options.EVMKeeper),
		circuitante.NewCircuitBreakerDecorator(options.CircuitKeeper),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		evmante.NewEVMAddressDecorator(options.EVMKeeper, options.EVMKeeper.AccountKeeper()), // TODO:
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
	}

	anteHandler := sdk.ChainAnteDecorators(anteDecorators...)

	evmAnteDecorators := []sdk.AnteDecorator{ // AnteFullDecorator
		evmante.NewEVMPreprocessDecorator(options.EVMKeeper, options.EVMKeeper.AccountKeeper(), EVMAssociatePriority),
		evmante.NewBasicDecorator(),
		evmante.NewEVMFeeCheckDecorator(options.EVMKeeper, EVMAssociatePriority), // TODO: is EVMAssociatePriority right?
		evmante.NewEVMSigVerifyDecorator(options.EVMKeeper, options.LatestCtxGetter),
		evmante.NewGasLimitDecorator(options.EVMKeeper),
	}
	evmAnteHandler := sdk.ChainAnteDecorators(evmAnteDecorators...)
	evmAnteHandler = evmante.NewAnteErrorHandler(evmAnteHandler, options.EVMKeeper).Handle

	router := evmante.NewEVMRouterDecorator(anteHandler, evmAnteHandler)

	return router.AnteHandle, nil
}
