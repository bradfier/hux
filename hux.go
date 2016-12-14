package hux

import (
	"fmt"
	"net/http"
)

type hux struct {
	baseURI     string
	accessToken string
}

func NewHux(baseURI, accessToken string) *hux {
	return &hux{baseURI, accessToken}
}

func (hux *hux) sendRequest(uri string) (resp *http.Response, err error) {
	uri = fmt.Sprintf("%s%s?accessToken=%s", hux.baseURI, uri, hux.accessToken)
	return http.Get(uri)
}
