package json

import (
	"encoding/xml"
)

type Request struct {
	CountryCode string `json:"countryCode"`
}

type CountriesRequest struct {
	XMLName        xml.Name `xml:"soapenv:Envelope"`
	XMLNamespace   string   `xml:"xmlns:soapenv,attr"`
	XMLHSNamespace string   `xml:"xmlns:hs,attr"`
	CountryCode    string   `xml:"soapenv:Body>hs:GetHolidaysAvailable>hs:countryCode"`
}

func (r Request) ToXML() CountriesRequest {
	return CountriesRequest{
		CountryCode:    r.CountryCode,
		XMLHSNamespace: "http://www.holidaywebservice.com/HolidayService_v2/",
		XMLNamespace:   "http://schemas.xmlsoap.org/soap/envelope/",
	}
}
