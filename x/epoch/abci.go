package epoch

import (
	"fmt"
	"github.com/provenance-io/provenance/x/epoch/keeper"
	"github.com/provenance-io/provenance/x/epoch/types"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker of epochs module
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	k.IterateEpochInfo(ctx, func(index int64, epochInfo types.EpochInfo) (stop bool) {
		logger := k.Logger(ctx)

		// Has it not started, and is the block height > initial epoch start height
		shouldInitialEpochStart := !epochInfo.EpochCountingStarted && !(epochInfo.StartHeight > ctx.BlockHeight())

		epochEndTime := epochInfo.CurrentEpochStartHeight + epochInfo.Duration
		shouldEpochStart := ctx.BlockHeight() > epochEndTime && !shouldInitialEpochStart && !(epochInfo.StartHeight > ctx.BlockHeight())

		if shouldInitialEpochStart || shouldEpochStart {
			epochInfo.CurrentEpochStartHeight = ctx.BlockHeight()

			if shouldInitialEpochStart {
				epochInfo.EpochCountingStarted = true
				epochInfo.CurrentEpoch = 1
				epochInfo.CurrentEpochStartHeight = epochInfo.StartHeight
				logger.Info(fmt.Sprintf("Starting new epoch with identifier %s epoch number %d", epochInfo.Identifier, epochInfo.CurrentEpoch))
			} else {
				epochInfo.CurrentEpoch += 1
				epochInfo.CurrentEpochStartHeight = epochInfo.CurrentEpochStartHeight + epochInfo.Duration
				logger.Info(fmt.Sprintf("Starting epoch with identifier %s epoch number %d", epochInfo.Identifier, epochInfo.CurrentEpoch))
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						types.EventTypeEpochEnd,
						sdk.NewAttribute(types.AttributeEpochNumber, fmt.Sprintf("%d", epochInfo.CurrentEpoch)),
					),
				)
				k.AfterEpochEnd(ctx, epochInfo.Identifier, epochInfo.CurrentEpoch)
			}
			k.SetEpochInfo(ctx, epochInfo)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeEpochStart,
					sdk.NewAttribute(types.AttributeEpochNumber, fmt.Sprintf("%d", epochInfo.CurrentEpoch)),
					sdk.NewAttribute(types.AttributeEpochStartHeight, fmt.Sprintf("%d", epochInfo.CurrentEpochStartHeight)),
				),
			)
			k.BeforeEpochStart(ctx, epochInfo.Identifier, epochInfo.CurrentEpoch)
		}

		return false
	})
}
