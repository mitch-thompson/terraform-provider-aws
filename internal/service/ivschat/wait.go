package ivschat

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ivschat"
	"github.com/aws/aws-sdk-go-v2/service/ivschat/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
)

func waitLoggingConfigurationCreated(ctx context.Context, conn *ivschat.Client, id string, timeout time.Duration) (*ivschat.GetLoggingConfigurationOutput, error) {
	stateConf := &resource.StateChangeConf{
		Pending:                   enum.Slice(types.LoggingConfigurationStateCreating),
		Target:                    enum.Slice(types.LoggingConfigurationStateActive),
		Refresh:                   statusLoggingConfiguration(ctx, conn, id),
		Timeout:                   timeout,
		NotFoundChecks:            20,
		ContinuousTargetOccurence: 2,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*ivschat.GetLoggingConfigurationOutput); ok {
		return out, err
	}

	return nil, err
}

func waitLoggingConfigurationUpdated(ctx context.Context, conn *ivschat.Client, id string, timeout time.Duration) (*ivschat.GetLoggingConfigurationOutput, error) {
	stateConf := &resource.StateChangeConf{
		Pending:                   enum.Slice(types.LoggingConfigurationStateUpdating),
		Target:                    enum.Slice(types.LoggingConfigurationStateActive),
		Refresh:                   statusLoggingConfiguration(ctx, conn, id),
		Timeout:                   timeout,
		NotFoundChecks:            20,
		ContinuousTargetOccurence: 2,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*ivschat.GetLoggingConfigurationOutput); ok {
		return out, err
	}

	return nil, err
}

func waitLoggingConfigurationDeleted(ctx context.Context, conn *ivschat.Client, id string, timeout time.Duration) (*ivschat.GetLoggingConfigurationOutput, error) {
	stateConf := &resource.StateChangeConf{
		Pending: enum.Slice(types.LoggingConfigurationStateDeleting, types.LoggingConfigurationStateActive),
		Target:  []string{},
		Refresh: statusLoggingConfiguration(ctx, conn, id),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*ivschat.GetLoggingConfigurationOutput); ok {
		return out, err
	}

	return nil, err
}
