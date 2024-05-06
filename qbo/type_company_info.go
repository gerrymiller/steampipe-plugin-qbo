package qbo

type ApiCompanyInfo struct {
	CompanyInfo CompanyInfo `json:"CompanyInfo"`
}

func (apiResponse *ApiCompanyInfo) GetResponse() *CompanyInfo {
	return &apiResponse.CompanyInfo
}

type CompanyInfo struct {
	ID                           string          `json:"Id"`
	SyncToken                    string          `json:"SyncToken"`
	CompanyName                  string          `json:"CompanyName"`
	CompanyAddress               PhysicalAddress `json:"CompanyAddr"`
	LegalAddress                 PhysicalAddress `json:"LegalAddr"`
	CustomerCommunicationAddress PhysicalAddress `json:"CustomerCommunicationAddr"`
	SupportedLanguages           string          `json:"SupportedLanguages"`
	Country                      string          `json:"Country"`
	Email                        EmailAddress    `json:"Email"`
	Web                          WebSiteAddress  `json:"WebAddr"`
	Attributes                   []NameValue     `json:"NameValue"`
	FiscalYearStartMonth         string          `json:"FiscalYearStartMonth"`
	PrimaryPhone                 TelephoneNumber `json:"PrimaryPhone"`
	LegalName                    string          `json:"LegalName"`
	MetaData                     MetaData        `json:"MetaData"`
	CompanyStartDate             string          `json:"CompanyStartDate"`
}
