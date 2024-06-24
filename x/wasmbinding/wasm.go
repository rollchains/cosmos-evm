package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	aclkeeper "github.com/cosmos/cosmos-sdk/x/accesscontrol/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	evmwasm "github.com/sei-protocol/sei-chain/x/evm/client/wasm"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
)

func RegisterCustomPlugins(
	// oracle *oraclekeeper.Keeper,
	// dex *dexkeeper.Keeper,
	// epoch *epochkeeper.Keeper,
	// tokenfactory *tokenfactorykeeper.Keeper,
	_ *authkeeper.AccountKeeper,
	router wasmkeeper.MessageRouter,
	channelKeeper wasmtypes.ChannelKeeper,
	capabilityKeeper wasmtypes.CapabilityKeeper,
	bankKeeper wasmtypes.Burner,
	unpacker codectypes.AnyUnpacker,
	portSource wasmtypes.ICS20TransferPortSource,
	aclKeeper aclkeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
) []wasmkeeper.Option {
	// dexHandler := dexwasm.NewDexWasmQueryHandler(dex)
	// oracleHandler := oraclewasm.NewOracleWasmQueryHandler(oracle)
	// epochHandler := epochwasm.NewEpochWasmQueryHandler(epoch)
	// tokenfactoryHandler := tokenfactorywasm.NewTokenFactoryWasmQueryHandler(tokenfactory)
	evmHandler := evmwasm.NewEVMQueryHandler(evmKeeper)
	wasmQueryPlugin := NewQueryPlugin(evmHandler)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerHandlerOpt := wasmkeeper.WithMessageHandler(
		CustomMessageHandler(router, channelKeeper, capabilityKeeper, bankKeeper, evmKeeper, unpacker, portSource),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerHandlerOpt,
	}
}
