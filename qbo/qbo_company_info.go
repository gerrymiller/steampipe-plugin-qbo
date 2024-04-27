
package qbo

import (
	"context"

 	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableQBOCompanyInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "qbo_company_info",
		Description: "Company Information from QuickBooks Online",
		List: &plugin.ListConfig{
			Hydrate: listQBOCompanyInfo,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the Company.",
				Transform:   transform.FromField("ID").NullIfZero(),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the Company.",
				Transform:   transform.FromField("Name").NullIfZero(),
			},
		},
	}
}


func listQBOCompanyInfo(
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (interface{}, error) {
	c := CompanyInfo { "123", "Gerry"}
	d.StreamListItem(ctx, c)
	return nil, nil
}

type CompanyInfo struct {
	ID string 
	Name string
}