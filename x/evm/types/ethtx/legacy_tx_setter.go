package ethtx

import "cosmossdk.io/math"

func (tx *LegacyTx) SetTo(v string) {
	tx.To = v
}

func (tx *LegacyTx) SetAmount(v math.Int) {
	tx.Amount = &v
}

func (tx *LegacyTx) SetGasPrice(v math.Int) {
	tx.GasPrice = &v
}
