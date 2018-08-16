package json2xml_test

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
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
func Test_ConvertJSON_Input_JSON_Should_Be_XML(t *testing.T) {
	expectedXML, _ := ioutil.ReadFile("./request.xml")
	var request Request
	jsonData := []byte(`{"countryCode":"UnitedStates"}`)
	json.Unmarshal(jsonData, &request)

	requestXML := request.ToXML()
	actualXML, _ := xml.MarshalIndent(requestXML, "", "\t")
	if string(expectedXML) != string(actualXML) {
		t.Errorf("expected \n%s \nbut it got \n%s", expectedXML, actualXML)
	}
}
