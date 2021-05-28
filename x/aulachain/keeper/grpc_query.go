package keeper

import (
	"github.com/infu/aulachain/x/aulachain/types"
)

var _ types.QueryServer = Keeper{}
