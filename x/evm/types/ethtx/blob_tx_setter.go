package ethtx

import "cosmossdk.io/math"

func (tx *BlobTx) SetTo(v string) {
	tx.To = v
}

func (tx *BlobTx) SetAmount(v math.Int) {
	tx.Amount = &v
}

func (tx *BlobTx) SetGasFeeCap(v math.Int) {
	tx.GasFeeCap = &v
}

func (tx *BlobTx) SetGasTipCap(v math.Int) {
	tx.GasTipCap = &v
}

func (tx *BlobTx) SetAccesses(v AccessList) {
	tx.Accesses = v
}

func (tx *BlobTx) SetBlobFeeCap(v math.Int) {
	tx.BlobFeeCap = &v
}

func (tx *BlobTx) SetBlobHashes(v [][]byte) {
	tx.BlobHashes = v
}

func (tx *BlobTx) SetBlobSidecar(v *BlobTxSidecar) {
	tx.Sidecar = v
}
