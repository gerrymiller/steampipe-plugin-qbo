
package qbo

import (
	"context"

 	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/rwestlund/quickbooks-go"
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
				Transform:   transform.FromField("CompanyName").NullIfZero(),
			},
		},
	}
}


func listQBOCompanyInfo(
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (interface{}, error) {
	config := GetConfig(d.Connection)

	token := &quickbooks.BearerToken {
		RefreshToken: *config.RefreshToken,
	}

	qbClient, err := quickbooks.NewQuickbooksClient(
		*config.ClientId,
		*config.ClientSecret,
		*config.RealmId,
		false,
		token)

	if err != nil {
		plugin.Logger(ctx).Error(err.Error())
		return nil, err
	}

	if qbClient == nil {
		plugin.Logger(ctx).Error("qbClient is nil")
		return nil, nil
	}


	token, err = qbClient.RefreshToken(token.RefreshToken)
		
	//var companyInfo CompanyInfo
	companyInfoRaw, err := qbClient.FetchCompanyInfo()
	//companyInfo = *companyInfoRaw

	if err != nil {
		plugin.Logger(ctx).Error(err.Error())
		return nil, err
	}

	d.StreamListItem(ctx, companyInfoRaw)

	//c := CompanyInfo { "123", "Gerry"}
	//d.StreamListItem(ctx, c)
	return nil, nil
}

/*
type CompanyInfo struct {
	ID string `json:"Id"`
	SyncToken string `json:"SyncToken"`
	CompanyName string `json:"CompanyName"`
	CompanyAddress PhysicalAddress `json:"CompanyAddr"`
	SupportedLanguages string `json:"SupportedLanguages"`
	Country string `json:"Country"`
	Email EmailAddress `json:"Email"`
	Web WebSiteAddress `json:"WebAddr"`
	Attributes []NameValue `json:"NameValue"`
	FiscalYearStartMonth string `json:"FiscalYearStartMonth"`
	CustomerCommunicationAddress PhysicalAddress `json:"CustomerCommunicationAddr"`
	PrimaryPhone TelephoneNumber `json:"PrimaryPhone"`
	LegalName string `json:"LegalName"`
	MetaData MetaData `json:"MetaData"`
	CompanyStartDate DateTime `json:"CompanyStartDate"`
}

type PhysicalAddress struct {
	ID string `json:"Id"`
	Line1 string `json:"Line1"`
	Line2 string `json:"Line2"`
	Line3 string `json:"Line3"`
	Line4 string `json:"Line4"`
	Line5 string `json:"Line5"`
	City string `json:"City"`
	Region string `json:"CountrySubDivisionCode"`
	PostalCode string `json:"PostalCode"`
	Country string `json:"Country"`
	Latitude string `json:"Lat"`
	Longitude string `json:"Long"`
}

type TelephoneNumber struct {
	FreeFormNumber string `json:"FreeFormNumber"`
}

type EmailAddress struct {
	Address string `json:"Address"`
}

type WebSiteAddress struct {
	URI string `json:"URI"`
}

type DateTime struct {
	DateTime string `json:"dateTime"`
}

type NameValue struct {
	Name string `json:"Name"`
	Value string `json:"Value"`
}

type MetaData struct {
	CreateTime string `json:"CreateTime"`
	LastUpdateTime string `json:"LastUpdateTime"`
}
*/