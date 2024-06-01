package qbo

type ApiReport struct {
	Report Report
}

func (apiResponse *ApiReport) GetResponse() *Report {
	return &apiResponse.Report
}

type Report struct {
	Header  Header  `json:"Header"`
	Rows    Rows    `json:"Rows"`
	Columns Columns `json:"Columns"`
}
