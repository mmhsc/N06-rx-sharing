package drugstore

import (
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/drugstore/internal/keeper"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/drugstore/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper = keeper.Keeper
)
