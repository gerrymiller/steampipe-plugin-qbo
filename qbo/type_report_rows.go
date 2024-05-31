package qbo

type Rows struct {
	Row []Row `json:"Row"`
}

type Row struct {
	Header  *RowHeader `json:"Header,omitempty"`
	Rows    *Rows      `json:"Rows,omitempty"`
	Type    string     `json:"type"`
	Group   string     `json:"group,omitempty"`
	Summary *Summary   `json:"Summary,omitempty"`
	ColData []ColData  `json:"ColData,omitempty"`
}

type RowHeader struct {
	ColData []ColData `json:"ColData"`
}

type ColData struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value"`
}

type Summary struct {
	ColData []ColData `json:"ColData"`
}
