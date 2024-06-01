package qbo

func (apiResponse *Report) GetResponse() *Report {
	return apiResponse
}

type Report struct {
	Header  Header  `json:"Header"`
	Rows    Rows    `json:"Rows"`
	Columns Columns `json:"Columns"`
}
