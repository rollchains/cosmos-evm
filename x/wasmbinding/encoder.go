package wasmbinding

import (
	"encoding/json"

	"cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evmwasm "github.com/sei-protocol/sei-chain/x/evm/client/wasm"
)

type SeiWasmMessage struct {
	PlaceOrders     json.RawMessage `json:"place_orders,omitempty"`
	CancelOrders    json.RawMessage `json:"cancel_orders,omitempty"`
	CreateDenom     json.RawMessage `json:"create_denom,omitempty"`
	MintTokens      json.RawMessage `json:"mint_tokens,omitempty"`
	BurnTokens      json.RawMessage `json:"burn_tokens,omitempty"`
	ChangeAdmin     json.RawMessage `json:"change_admin,omitempty"`
	SetMetadata     json.RawMessage `json:"set_metadata,omitempty"`
	CallEVM         json.RawMessage `json:"call_evm,omitempty"`
	DelegateCallEVM json.RawMessage `json:"delegate_call_evm,omitempty"`
}

// TODO: this requires a wasm fork
// func CustomEncoder(sender sdk.AccAddress, msg json.RawMessage, info wasmvmtypes.MessageInfo, codeInfo wasmtypes.CodeInfo) ([]sdk.Msg, error) {
func CustomEncoder(sender sdk.AccAddress, msg json.RawMessage) ([]sdk.Msg, error) {
	var parsedMessage SeiWasmMessage
	if err := json.Unmarshal(msg, &parsedMessage); err != nil {
		return []sdk.Msg{}, errors.Wrap(err, "Error parsing Sei Wasm Message")
	}
	switch {
	case parsedMessage.CallEVM != nil:
		return evmwasm.EncodeCallEVM(parsedMessage.CallEVM, sender)
	case parsedMessage.DelegateCallEVM != nil:
		// return evmwasm.EncodeDelegateCallEVM(parsedMessage.DelegateCallEVM, sender, info, codeInfo)
		return []sdk.Msg{}, errors.Wrap(sdkerrors.ErrUnknownRequest, "parsedMessage.DelegateCallEVM not implemented yet (requires wasmd fork?? idk)")
	default:
		return []sdk.Msg{}, wasmvmtypes.UnsupportedRequest{Kind: "Unknown Sei Wasm Message"}
	}
}
