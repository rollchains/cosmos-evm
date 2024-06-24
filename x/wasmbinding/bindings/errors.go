package bindings

import (
	sdkErrors "cosmossdk.io/errors"
)

// Codes for wasm contract errors
var (
	DefaultCodespace = "wasmbinding"

	ErrParsingSeiWasmMsg = sdkErrors.Register(DefaultCodespace, 2, "Error parsing Sei Wasm Message")
)
