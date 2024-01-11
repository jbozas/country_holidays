package model

type Holiday struct {
	Date        string   `json:"date"`
	LocalName   string   `json:"localName"`
	CountryCode string   `json:"countryCode"`
	Name        string   `json:"name"`
	Global      bool     `json:"global"`
	Type        []string `json:"types"`
}