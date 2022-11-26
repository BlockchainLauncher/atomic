package keeper

import (
	"money/x/money/types"
)

var _ types.QueryServer = Keeper{}
