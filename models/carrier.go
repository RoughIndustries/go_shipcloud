package models

type Carriers struct {
	Carriers []Carrier
}

type Carrier struct {
	Name               string   `json:"name"`
	DisplayName        string   `json:"display_name"`
	Services           []string `json:"services"`
	PackageTypes       []string `json:"package_types"`
	AdditionalServices []string `json:"additional_services"`
	LabelFormats       struct {
		OneDay                    []string `json:"one_day"`
		UpsExpress1200            []string `json:"ups_express_1200"`
		Standard                  []string `json:"standard"`
		GlsExpress1200            []string `json:"gls_express_1200"`
		Returns                   []string `json:"returns"`
		DpagWarenpostSignature    []string `json:"dpag_warenpost_signature"`
		DhlPrio                   []string `json:"dhl_prio"`
		OneDayEarly               []string `json:"one_day_early"`
		DpagWarenpostUntracked    []string `json:"dpag_warenpost_untracked"`
		DhlEuropaket              []string `json:"dhl_europaket"`
		CargoInternationalExpress []string `json:"cargo_international_express"`
		DpagWarenpost             []string `json:"dpag_warenpost"`
		GlsExpress0900            []string `json:"gls_express_0900"`
		GlsExpress0800            []string `json:"gls_express_0800"`
		DhlWarenpost              []string `json:"dhl_warenpost"`
		GlsExpress1000            []string `json:"gls_express_1000"`
	} `json:"label_formats"`
}
