package hux

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestJSONUnmarshal(t *testing.T) {
	dat, err := ioutil.ReadFile("./test_service_details.json")

	if err != nil {
		t.Fatal("Failed to open JSON file for testing.")
	}

	var sd ServiceDetails

	if err := json.Unmarshal(dat, &sd); err != nil {
		t.Errorf("Failed to unmarshal JSON: {}", err)
	}

	if len(sd.SubsequentCallingPoints[0].CallingPoint) != 3 ||
		len(sd.PreviousCallingPoints[0].CallingPoint) != 21 {
		t.Error("Parsed incorrect number of calling points from JSON")
	}
}
