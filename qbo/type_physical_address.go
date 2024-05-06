package qbo

type PhysicalAddress struct {
	ID         string `json:"Id"`
	Line1      string `json:"Line1"`
	Line2      string `json:"Line2"`
	Line3      string `json:"Line3"`
	Line4      string `json:"Line4"`
	Line5      string `json:"Line5"`
	City       string `json:"City"`
	Region     string `json:"CountrySubDivisionCode"`
	PostalCode string `json:"PostalCode"`
	Country    string `json:"Country"`
	Latitude   string `json:"Lat"`
	Longitude  string `json:"Long"`
}
