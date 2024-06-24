package ethtx

import "cosmossdk.io/math"

func (tx *DynamicFeeTx) SetTo(v string) {
	tx.To = v
}

func (tx *DynamicFeeTx) SetAmount(v math.Int) {
	tx.Amount = &v
}

func (tx *DynamicFeeTx) SetGasFeeCap(v math.Int) {
	tx.GasFeeCap = &v
}

func (tx *DynamicFeeTx) SetGasTipCap(v math.Int) {
	tx.GasTipCap = &v
}

func (tx *DynamicFeeTx) SetAccesses(v AccessList) {
	tx.Accesses = v
}
