package hux

import (
	"errors"
	"os"
	"testing"
)

func getTestEnvironement() (accessToken string, baseUri string, err error) {
	accessToken = os.Getenv("HUX_ACCESS_TOKEN")
	baseUri = os.Getenv("HUX_URI")

	if accessToken == "" || baseUri == "" {
		err = errors.New("HUX_ACCESS_TOKEN and HUX_URI must be defined to test requests.")
	}
	return
}

func TestHTTPRequest(t *testing.T) {
	accessToken, baseUri, err := getTestEnvironement()

	if err != nil {
		t.Skip(err)
	}

	hux := NewHux(baseUri, accessToken)
	_, err = hux.sendRequest("/all/KGX")

	if err != nil {
		t.Fatalf("Failed to make HTTP request: %s", err)
	}
}

func TestJSONDecoder(t *testing.T) {
	accessToken, baseUri, err := getTestEnvironement()

	if err != nil {
		t.Skip(err)
	}

	hux := NewHux(baseUri, accessToken)
	crs, err := hux.GetCRSCodes("Kings Cross")

	if err != nil {
		t.Fatal(err)
	}

	if len(*crs) < 1 {
		t.Fatal("Didn't get CRS code for station.")
	}
}
