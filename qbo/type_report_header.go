package qbo

import "time"

type Header struct {
	Customer           string    `json:"Customer"`
	ReportName         string    `json:"ReportName"`
	Option             []Option  `json:"Option"`
	ReportBasis        string    `json:"ReportBasis"`
	StartPeriod        string    `json:"StartPeriod"`
	Currency           string    `json:"Currency"`
	EndPeriod          string    `json:"EndPeriod"`
	Time               time.Time `json:"Time"`
	SummarizeColumnsBy string    `json:"SummarizeColumnsBy"`
}
