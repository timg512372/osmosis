package stableswap

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

var _ types.PoolI = &Pool{}

func (pa Pool) GetAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(pa.Address)
	if err != nil {
		panic(fmt.Sprintf("could not bech32 decode address of pool with id: %d", pa.GetId()))
	}
	return addr
}

func (pa Pool) String() string {
	out, err := json.Marshal(pa)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (pa Pool) GetId() uint64 {
	return pa.Id
}

func (pa Pool) GetSwapFee(ctx sdk.Context) sdk.Dec {
	return pa.PoolParams.SwapFee
}
func (pa Pool) GetExitFee(ctx sdk.Context) sdk.Dec {
	return pa.PoolParams.ExitFee
}

func (pa Pool) IsActive(ctx sdk.Context) bool {
	return true
}

// Returns the coins in the pool owned by all LP shareholders
func (pa Pool) GetTotalPoolLiquidity(ctx sdk.Context) sdk.Coins {
	return pa.PoolLiquidity
}

func (pa Pool) GetTotalShares() sdk.Int {
	return pa.TotalShares.Amount
}

// These should all get moved to amm.go
func (pa Pool) CalcOutAmtGivenIn(ctx sdk.Context, tokenIn sdk.Coins, tokenOutDenom string, swapFee sdk.Dec) (tokenOut sdk.DecCoin, err error) {
	return sdk.DecCoin{}, types.ErrNotImplemented
}
func (pa *Pool) SwapOutAmtGivenIn(ctx sdk.Context, tokenIn sdk.Coins, tokenOutDenom string, swapFee sdk.Dec) (tokenOut sdk.Coin, err error) {
	return sdk.Coin{}, types.ErrNotImplemented
}

func (pa Pool) CalcInAmtGivenOut(ctx sdk.Context, tokenOut sdk.Coins, tokenInDenom string, swapFee sdk.Dec) (tokenIn sdk.DecCoin, err error) {
	return sdk.DecCoin{}, types.ErrNotImplemented
}
func (pa *Pool) SwapInAmtGivenOut(ctx sdk.Context, tokenOut sdk.Coins, tokenInDenom string, swapFee sdk.Dec) (tokenIn sdk.Coin, err error) {
	return sdk.Coin{}, types.ErrNotImplemented
}

func (pa Pool) SpotPrice(ctx sdk.Context, baseAssetDenom string, quoteAssetDenom string) (sdk.Dec, error) {
	return sdk.Dec{}, types.ErrNotImplemented
}

func (pa Pool) CalcJoinPoolShares(ctx sdk.Context, tokensIn sdk.Coins, swapFee sdk.Dec) (numShares sdk.Int, newLiquidity sdk.Coins, err error) {
	return sdk.Int{}, sdk.Coins{}, types.ErrNotImplemented
}
func (pa *Pool) JoinPool(ctx sdk.Context, tokensIn sdk.Coins, swapFee sdk.Dec) (numShares sdk.Int, err error) {
	return sdk.Int{}, types.ErrNotImplemented
}

func (pa *Pool) ExitPool(ctx sdk.Context, numShares sdk.Int, exitFee sdk.Dec) (exitedCoins sdk.Coins, err error) {
	return sdk.Coins{}, types.ErrNotImplemented
}

func (pa Pool) CalcExitPoolShares(ctx sdk.Context, numShares sdk.Int, exitFee sdk.Dec) (exitedCoins sdk.Coins, err error) {
	return sdk.Coins{}, types.ErrNotImplemented
}

func (pa *Pool) PokePool(blockTime time.Time) {
	panic("not implemented yet")
}
