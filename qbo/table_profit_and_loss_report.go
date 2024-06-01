package qbo

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableQBOProfitAndLossReport(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "qbo_profit_and_loss_report",
		Description: "Profit and Loss Report from QuickBooks Online",
		List: &plugin.ListConfig{
			Hydrate: listQBOProfitAndLossReport,
		},
		Columns: []*plugin.Column{
			{
				Name:        "header",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Transform:   transform.FromField("Header").NullIfZero(),
			},
			{
				Name:        "rows",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Transform:   transform.FromField("Rows").NullIfZero(),
			},
			{
				Name:        "columns",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Transform:   transform.FromField("Columns").NullIfZero(),
			},
		},
	}
}

func listQBOProfitAndLossReport(
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (interface{}, error) {

	reportApi := new(ApiReport)
	_, err := qboApiCall(&reportApi,
		"{{.baseURL}}/v3/company/{{.realmId}}/reports/ProfitAndLoss",
		ctx,
		d,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get profit and loss report: %v", err)
	}

	plugin.Logger(ctx).Info("Profit and Loss Report: ", reportApi.GetResponse())
	d.StreamListItem(ctx, *reportApi.GetResponse())
	return nil, nil
}
