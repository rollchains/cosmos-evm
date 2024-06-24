package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	evmwasm "github.com/sei-protocol/sei-chain/x/evm/client/wasm"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
)

func RegisterCustomPlugins(
	c codec.Codec,
	ics4Wrapper wasmtypes.ICS4Wrapper,
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
		CustomMessageHandler(c, router, ics4Wrapper, channelKeeper, capabilityKeeper, bankKeeper, evmKeeper, unpacker, portSource),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerHandlerOpt,
	}
}
