package etf

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/defund-labs/defund/x/etf/keeper"
	"github.com/defund-labs/defund/x/etf/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the fund
	for _, elem := range genState.FundList {
		k.SetFund(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.FundList = k.GetAllFund(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
