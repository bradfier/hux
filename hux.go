package hux

import (
	"encoding/json"
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

func (hux *hux) GetCRSCodes(filter string) (stationCodes *[]CRSStationCode, err error) {
	uri := "/crs/" + filter
	resp, err := hux.sendRequest(uri)

	if err != nil {
		return stationCodes, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&stationCodes); err != nil {
		return stationCodes, err
	}

	return stationCodes, err
}

func (hux *hux) GetAllCRSCodes() (stationCodes *[]CRSStationCode, err error) {
	return hux.GetCRSCodes("")
}

func (hux *hux) GetDepartures(hq huxQuery) (ts *boardResponse, err error) {
	uri := fmt.Sprintf("/departures/%s", hq)

	resp, err := hux.sendRequest(uri)

	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&ts)

	return ts, err
}
