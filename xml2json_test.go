package json2xml_test

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

type GetCountriesAvailable struct {
	CountryCode []CountryCode `xml:"Body>GetCountriesAvailableResponse>GetCountriesAvailableResult>CountryCode" json:"countries"`
}

type CountryCode struct {
	Code        string `xml:"Code" json:"code"`
	Description string `xml:"Description" json:"description"`
}

func Test_ConvertXML_Input_XML_Should_Be_Struct(t *testing.T) {
	var actual GetCountriesAvailable
	xmlFile, _ := ioutil.ReadFile("./response.xml")
	expected := GetCountriesAvailable{
		CountryCode: []CountryCode{
			CountryCode{"Canada", "Canada"},
			CountryCode{"GreatBritain", "Great Britain and Wales"},
			CountryCode{"IrelandNorthern", "Northern Ireland"},
			CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
			CountryCode{"Scotland", "Scotland"},
			CountryCode{"UnitedStates", "United States"},
		},
	}

	xml.Unmarshal(xmlFile, &actual)

	for index, actualCountryCode := range actual.CountryCode {
		expectedCountryCode := expected.CountryCode[index]
		if actualCountryCode != expectedCountryCode {
			t.Errorf("expected at index: %d %s but it got %s", index, expectedCountryCode, actualCountryCode)
		}
	}
}

func Test_ToJSON_Should_Be_JSON(t *testing.T) {
	expected := `{"countries":[{"code":"Canada","description":"Canada"},{"code":"GreatBritain","description":"Great Britain and Wales"},{"code":"IrelandNorthern","description":"Northern Ireland"},{"code":"IrelandRepublicOf","description":"Republic of Ireland"},{"code":"Scotland","description":"Scotland"},{"code":"UnitedStates","description":"United States"}]}`
	getCountriesAvailable := GetCountriesAvailable{
		CountryCode: []CountryCode{
			CountryCode{"Canada", "Canada"},
			CountryCode{"GreatBritain", "Great Britain and Wales"},
			CountryCode{"IrelandNorthern", "Northern Ireland"},
			CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
			CountryCode{"Scotland", "Scotland"},
			CountryCode{"UnitedStates", "United States"},
		},
	}
	actual, _ := json.Marshal(getCountriesAvailable)
	if expected != string(actual) {
		t.Errorf("expected \n%s but it got \n%s", expected, actual)
	}
}
