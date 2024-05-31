package qbo

type Columns struct {
	Column []Column `json:"Column"`
}

type Column struct {
	ColType  string     `json:"ColType"`
	ColTitle string     `json:"ColTitle"`
	MetaData []MetaData `json:"MetaData"`
}
