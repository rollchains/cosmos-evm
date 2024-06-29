package keeper

import (
	"cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func (k *Keeper) AddAnteSurplus(ctx sdk.Context, txHash common.Hash, surplus math.Int) error {
	store := prefix.NewStore(ctx.KVStore(k.memStoreKey), types.AnteSurplusPrefix)
	bz, err := surplus.Marshal()
	if err != nil {
		return err
	}
	store.Set(txHash[:], bz)
	return nil
}

func (k *Keeper) GetAnteSurplusSum(ctx sdk.Context) math.Int {
	iter := prefix.NewStore(ctx.KVStore(k.memStoreKey), types.AnteSurplusPrefix).Iterator(nil, nil)
	defer iter.Close()
	res := math.ZeroInt()
	for ; iter.Valid(); iter.Next() {
		surplus := math.Int{}
		_ = surplus.Unmarshal(iter.Value())
		res = res.Add(surplus)
	}
	return res
}

func (k *Keeper) DeleteAllAnteSurplus(ctx sdk.Context) {
	// TODO: old: _ = prefix.NewStore(ctx.KVStore(k.memStoreKey), types.AnteSurplusPrefix).DeleteAll(nil, nil)
	s := prefix.NewStore(ctx.KVStore(k.memStoreKey), types.AnteSurplusPrefix)
	iter := s.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		s.Delete(iter.Key())
	}

}
