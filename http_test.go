package hux

import (
	"os"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	accessToken := os.Getenv("HUX_ACCESS_TOKEN")
	baseUri := os.Getenv("HUX_URI")

	if accessToken == "" || baseUri == "" {
		t.Skipf("HUX_ACCESS_TOKEN and HUX_URI must be defined to test requests.")
	}

	hux := NewHux(baseUri, accessToken)

	_, err := hux.sendRequest("/all/KGX")

	if err != nil {
		t.Fatalf("Failed to make HTTP request: %s", err)
	}
}
