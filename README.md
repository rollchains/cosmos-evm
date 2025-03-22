> This repo is now archived due to evm-os becoming open source apache2: https://github.com/cosmos/evm - no further developments to occur here.

# Sei EVM

A v0.50 impl of the sei network EVM. Following the same repo structure, just updated to latest

Built off commit: f6137af10683b8dd5f06b6f07242f9f713b0c2ff (Mon Jun 24 12:21:25 2024 +0800) from sei-chain.

## Steps
- cp evmrpc
- cd evmrpc find . -type f -name '*_test.go' -exec sh -c 'x="{}"; mv "$x" "${x}.archive"' \;
- copy in utils/ (no breaks)
- copy in precompiles (rename test)
- copy in x/evm, rename files
- replace github.com/cometbft/cometbft -> github.com/cometbft/cometbft
- evm ante decorators modified to not use constants, but allow NewAnte setup for EVMAssociatePriority & MaxPriority vars


# other
- question: wasm bindings required? maybe we don't allow cw20 pairings? but does this break ERC20<->bank? we do not care about ERC20<->CW. (x/evm/client/wasm/encoder.go)
- are tokenfactory bindings required? (so ERC20 / CW20 can be minted and handled? or does it use native bank for the pointer contract logic)
