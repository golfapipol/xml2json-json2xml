package json_test

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
	. "xml2json-json2xml/json"
)

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
