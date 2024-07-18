TODO:

ProcessBlock has an evm thing. Is this an SDK ext, do we ned?

func (app *App) ProcessBlock(ctx sdk.Context, txs [][]byte, req BlockProcessRequest, lastCommit abci.CommitInfo) ([]abci.Event, []*abci.ExecTxResult, abci.ResponseEndBlock, error) {

        ...

    	app.EvmKeeper.SetTxResults(txResults)

}