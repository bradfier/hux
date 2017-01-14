package hux

import (
	"errors"
	"os"
	"testing"
)

func getTestEnvironment() (accessToken string, baseURI string, err error) {
	accessToken = os.Getenv("HUX_ACCESS_TOKEN")
	baseURI = os.Getenv("HUX_URI")

	if accessToken == "" || baseURI == "" {
		err = errors.New("HUX_ACCESS_TOKEN and HUX_URI must be defined to test requests")
	}
	return
}

func TestHTTPRequest(t *testing.T) {
	accessToken, baseURI, err := getTestEnvironment()

	if err != nil {
		t.Skip(err)
	}

	hux := NewHux(baseURI, accessToken)
	_, err = hux.sendRequest("/all/KGX")

	if err != nil {
		t.Fatalf("Failed to make HTTP request: %s", err)
	}
}

func TestJSONDecoder(t *testing.T) {
	accessToken, baseURI, err := getTestEnvironment()

	if err != nil {
		t.Skip(err)
	}

	hux := NewHux(baseURI, accessToken)
	crs, err := hux.GetCRSCodes("Kings Cross")

	if err != nil {
		t.Fatal(err)
	}

	if len(*crs) < 1 {
		t.Fatal("Didn't get CRS code for station.")
	}
}

func TestGetDeparture(t *testing.T) {
	accessToken, baseURI, err := getTestEnvironment()

	if err != nil {
		t.Skip(err)
	}

	hux := NewHux(baseURI, accessToken)
	hq := new(HuxQueryBuilder).QueryStation("KGX").Build()

	ts, err := hux.GetDepartures(hq)

	if err != nil {
		t.Fatal(err)
	}

	if ts == nil {
		t.Error("no departures decoded")
	}
}
