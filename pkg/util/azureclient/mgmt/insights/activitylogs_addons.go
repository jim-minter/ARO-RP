package insights

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	mgmtinsights "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2018-03-01/insights"
)

// ActivityLogsClientAddons contains addons for ActivityLogsClient
type ActivityLogsClientAddons interface {
	List(ctx context.Context, filter string, selectParameter string) (result []mgmtinsights.EventData, err error)
}

func (c *activityLogsClient) List(ctx context.Context, filter string, selectParameter string) (result []mgmtinsights.EventData, err error) {
	page, err := c.ActivityLogsClient.List(ctx, filter, selectParameter)
	if err != nil {
		return nil, err
	}

	for page.NotDone() {
		result = append(result, page.Values()...)

		err = page.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
