package qbo

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
				Name:        "customer",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "report_name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "start_period",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "end_period",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "group",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "col_title",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
		},
	}
}

func createColumnMap(columns Columns) map[string]string {
	columnMap := make(map[string]string)
	for _, column := range columns.Column {
		for _, meta := range column.MetaData {
			if meta.Name == "ColKey" {
				columnMap[meta.Value] = column.ColTitle
			}
		}
	}
	return columnMap
}

func extractRows(rows []Row, parentGroup string, header Header, columnMap map[string]string) []map[string]interface{} {
	var result []map[string]interface{}
	for _, row := range rows {
		group := row.Group
		if group == "" {
			group = parentGroup
		}

		if row.Header != nil {
			for _, col := range row.Header.ColData {
				colTitle := columnMap[col.ID]
				result = append(result, map[string]interface{}{
					"customer":     header.Customer,
					"report_name":  header.ReportName,
					"start_period": header.StartPeriod,
					"end_period":   header.EndPeriod,
					"group":        group,
					"type":         row.Type,
					"col_title":    colTitle,
					"value":        col.Value,
				})
			}
		}

		if row.Rows != nil {
			result = append(result, extractRows(row.Rows.Row, group, header, columnMap)...)
		}

		if row.Summary != nil {
			for _, col := range row.Summary.ColData {
				colTitle := columnMap[col.ID]
				result = append(result, map[string]interface{}{
					"customer":     header.Customer,
					"report_name":  header.ReportName,
					"start_period": header.StartPeriod,
					"end_period":   header.EndPeriod,
					"group":        group,
					"type":         row.Type,
					"col_title":    colTitle,
					"value":        col.Value,
				})
			}
		}

		for _, col := range row.ColData {
			colTitle := columnMap[col.ID]
			result = append(result, map[string]interface{}{
				"customer":     header.Customer,
				"report_name":  header.ReportName,
				"start_period": header.StartPeriod,
				"end_period":   header.EndPeriod,
				"group":        group,
				"type":         row.Type,
				"col_title":    colTitle,
				"value":        col.Value,
			})
		}
	}
	return result
}

func listQBOProfitAndLossReport(
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (interface{}, error) {

	reportApi := new(Report)
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

	columnMap := createColumnMap(reportApi.GetResponse().Columns)

	rows := extractRows(reportApi.GetResponse().Rows.Row, "", reportApi.GetResponse().Header, columnMap)
	for _, row := range rows {
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}
