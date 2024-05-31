package qbo

type Report struct {
	Header  Header  `json:"Header"`
	Rows    Rows    `json:"Rows"`
	Columns Columns `json:"Columns"`
}
