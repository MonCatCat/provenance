package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/provenance-io/provenance/x/exchange"
)

func CmdQuery() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        exchange.ModuleName,
		Aliases:                    []string{"ex"},
		Short:                      "Querying commands for the exchange module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryOrderFeeCalc(),
		CmdQueryGetOrder(),
		CmdQueryGetOrderByExternalID(),
		CmdQueryGetMarketOrders(),
		CmdQueryGetOwnerOrders(),
		CmdQueryGetAssetOrders(),
		CmdQueryGetAllOrders(),
		CmdQueryGetMarket(),
		CmdQueryGetAllMarkets(),
		CmdQueryParams(),
		CmdQueryValidateCreateMarket(),
		CmdQueryValidateMarket(),
		CmdQueryValidateManageFees(),
	)

	return cmd
}

// CmdQueryOrderFeeCalc creates the order-fee-calc sub-command for the exchange query command.
func CmdQueryOrderFeeCalc() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "order-fee-calc",
		Aliases: []string{"fee-calc", "order-calc"},
		Short:   "Calculate the fees for an order",
		RunE:    genericQueryRunE(MakeQueryOrderFeeCalc, exchange.QueryClient.OrderFeeCalc),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryOrderFeeCalc(cmd)
	return cmd
}

// CmdQueryGetOrder creates the order sub-command for the exchange query command.
func CmdQueryGetOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "order",
		Aliases: []string{"get-order"},
		Short:   "Get an order by id",
		RunE:    genericQueryRunE(MakeQueryGetOrder, exchange.QueryClient.GetOrder),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetOrder(cmd)
	return cmd
}

// CmdQueryGetOrderByExternalID creates the order-by-external-id sub-command for the exchange query command.
func CmdQueryGetOrderByExternalID() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "order-by-external-id",
		Aliases: []string{"get-order-by-external-id", "by-external-id", "external-id"},
		Short:   "Get an order by market id and external id",
		RunE:    genericQueryRunE(MakeQueryGetOrderByExternalID, exchange.QueryClient.GetOrderByExternalID),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetOrderByExternalID(cmd)
	return cmd
}

// CmdQueryGetMarketOrders creates the market-orders sub-command for the exchange query command.
func CmdQueryGetMarketOrders() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "market-orders",
		Aliases: []string{"get-market-orders"},
		Short:   "Look up orders for a market",
		RunE:    genericQueryRunE(MakeQueryGetMarketOrders, exchange.QueryClient.GetMarketOrders),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetMarketOrders(cmd)
	return cmd
}

// CmdQueryGetOwnerOrders creates the owner-orders sub-command for the exchange query command.
func CmdQueryGetOwnerOrders() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "owner-orders",
		Aliases: []string{"get-owner-orders"},
		Short:   "Look up orders with a specific owner",
		RunE:    genericQueryRunE(MakeQueryGetOwnerOrders, exchange.QueryClient.GetOwnerOrders),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetOwnerOrders(cmd)
	return cmd
}

// CmdQueryGetAssetOrders creates the asset-orders sub-command for the exchange query command.
func CmdQueryGetAssetOrders() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "asset-orders",
		Aliases: []string{"get-asset-orders", "denom-orders", "get-denom-orders", "assets-orders", "get-assets-orders"},
		Short:   "Look up orders with a specific asset denom",
		RunE:    genericQueryRunE(MakeQueryGetAssetOrders, exchange.QueryClient.GetAssetOrders),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetAssetOrders(cmd)
	return cmd
}

// CmdQueryGetAllOrders creates the all-orders sub-command for the exchange query command.
func CmdQueryGetAllOrders() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "all-orders",
		Aliases: []string{"get-all-orders"},
		Short:   "Get all orders",
		RunE:    genericQueryRunE(MakeQueryGetAllOrders, exchange.QueryClient.GetAllOrders),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetAllOrders(cmd)
	return cmd
}

// CmdQueryGetMarket creates the market sub-command for the exchange query command.
func CmdQueryGetMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "market",
		Aliases: []string{"get-market"},
		Short:   "Get market setup and information",
		RunE:    genericQueryRunE(MakeQueryGetMarket, exchange.QueryClient.GetMarket),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetMarket(cmd)
	return cmd
}

// CmdQueryGetAllMarkets creates the all-markets sub-command for the exchange query command.
func CmdQueryGetAllMarkets() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "all-markets",
		Aliases: []string{"get-all-markets"},
		Short:   "Get all markets",
		RunE:    genericQueryRunE(MakeQueryGetAllMarkets, exchange.QueryClient.GetAllMarkets),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryGetAllMarkets(cmd)
	return cmd
}

// CmdQueryParams creates the params sub-command for the exchange query command.
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Aliases: []string{"get-params"},
		Short:   "Get the exchange module params",
		RunE:    genericQueryRunE(MakeQueryParams, exchange.QueryClient.Params),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryParams(cmd)
	return cmd
}

// CmdQueryValidateCreateMarket creates the validate-create-market sub-command for the exchange query command.
func CmdQueryValidateCreateMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "validate-create-market",
		Aliases: []string{"create-market-validate"},
		Short:   "Validate a create market request",
		RunE:    genericQueryRunE(MakeQueryValidateCreateMarket, exchange.QueryClient.ValidateCreateMarket),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryValidateCreateMarket(cmd)
	return cmd
}

// CmdQueryValidateMarket creates the validate-market sub-command for the exchange query command.
func CmdQueryValidateMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "validate-market",
		Aliases: []string{"market-validate"},
		Short:   "Validate an existing market's setup",
		RunE:    genericQueryRunE(MakeQueryValidateMarket, exchange.QueryClient.ValidateMarket),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryValidateMarket(cmd)
	return cmd
}

// CmdQueryValidateManageFees creates the validate-manage-fees sub-command for the exchange query command.
func CmdQueryValidateManageFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "validate-manage-fees",
		Aliases: []string{"manage-fees-validate"},
		Short:   "Validate a manage fees request",
		RunE:    genericQueryRunE(MakeQueryValidateManageFees, exchange.QueryClient.ValidateManageFees),
	}

	flags.AddQueryFlagsToCmd(cmd)
	SetupCmdQueryValidateManageFees(cmd)
	return cmd
}
