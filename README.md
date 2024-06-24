# Sei EVM

A v0.50 impl of the sei network EVM

Built off commit: f6137af10683b8dd5f06b6f07242f9f713b0c2ff (Mon Jun 24 12:21:25 2024 +0800) from sei-chain.

## Steps
- cp evmrpc
- cd evmrpc find . -type f -name '*_test.go' -exec sh -c 'x="{}"; mv "$x" "${x}.archive"' \;
- copy in utils/ (no breaks)
- copy in precompiles (rename test)
- replace github.com/cometbft/cometbft -> github.com/cometbft/cometbft