TODO:

ProcessBlock has an evm thing. Is this an SDK ext, do we ned?

func (app *App) ProcessBlock(ctx sdk.Context, txs [][]byte, req BlockProcessRequest, lastCommit abci.CommitInfo) ([]abci.Event, []*abci.ExecTxResult, abci.ResponseEndBlock, error) {

        ...

    	app.EvmKeeper.SetTxResults(txResults)

}


NOTE: to compile I edit `../../../.gvm/pkgsets/go1.22.5/global/pkg/mod/github.com/sei-protocol/go-ethereum@v1.13.5-sei-22/ethdb/pebble/pebble.go:592:21` atm to also have an error
I tried to go mod why this & can not narrow down atm. Will fix later

TODO:
add cosmos.msg.v1.service to proto messages