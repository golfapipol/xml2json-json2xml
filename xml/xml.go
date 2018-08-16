package xml

// encoding/json
type CountriesResponse struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// encoding/xml
type GetCountriesAvailable struct {
	CountryCode []CountryCode `xml:"Body>GetCountriesAvailableResponse>GetCountriesAvailableResult>CountryCode"`
}

type CountryCode struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

func (g GetCountriesAvailable) ToJSON() CountriesResponse {
	countries := make([]Country, len(g.CountryCode))
	for index := range g.CountryCode {
		countries[index] = Country{
			Code:        g.CountryCode[index].Code,
			Description: g.CountryCode[index].Description,
		}
	}
	return CountriesResponse{
		Countries: countries,
	}
}
