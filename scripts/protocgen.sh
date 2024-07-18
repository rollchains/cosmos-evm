set -eo pipefail

echo "Generating gogo proto code"
cd proto
buf mod update
cd ..
buf generate

# move proto files to the right places
cp -r ./github.com/sei-protocol/sei-chain/x/* x/
rm -rf ./github.com

