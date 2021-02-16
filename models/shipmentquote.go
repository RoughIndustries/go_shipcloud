package models

type ShipmentQuoteInput struct {
	Carrier string `json:"carrier"`
	Service string `json:"service"`
	To      struct {
		Street   string `json:"street"`
		StreetNo string `json:"street_no"`
		City     string `json:"city"`
		ZipCode  string `json:"zip_code"`
		Country  string `json:"country"`
	} `json:"to"`
	Package struct {
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
		Length float64 `json:"length"`
		Weight float64 `json:"weight"`
		Type   string  `json:"type"`
	} `json:"package"`
	From struct {
		Street   string `json:"street"`
		StreetNo string `json:"street_no"`
		City     string `json:"city"`
		ZipCode  string `json:"zip_code"`
		Country  string `json:"country"`
	} `json:"from"`
}

type ShipmentQuoteResult struct {
	ShipmentQuote struct {
		Price float64 `json:"price"`
	} `json:"shipment_quote"`
}
