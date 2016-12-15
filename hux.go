package hux

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hux struct {
	baseURI     string
	accessToken string
}

func NewHux(baseURI, accessToken string) *Hux {
	return &Hux{baseURI, accessToken}
}

func (hux *Hux) sendRequest(uri string) (resp *http.Response, err error) {
	uri = fmt.Sprintf("%s%s?accessToken=%s", hux.baseURI, uri, hux.accessToken)
	return http.Get(uri)
}

func (hux *Hux) doRequest(uri string, data interface{}) error {
	resp, err := hux.sendRequest(uri)

	if err != nil {
		return err
	}

	if err = json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return err
}

func (hux *Hux) GetCRSCodes(filter string) (stationCodes *[]CRSStationCode, err error) {
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

func (hux *Hux) GetAllCRSCodes() (stationCodes *[]CRSStationCode, err error) {
	return hux.GetCRSCodes("")
}

func (hux *Hux) GetDepartures(hq huxQuery) (ts *boardResponse, err error) {
	ts = new(boardResponse)
	uri := fmt.Sprintf("/departures/%s", hq)
	err = hux.doRequest(uri, ts)
	return
}

func (hux *Hux) GetArrivals(hq huxQuery) (ts *boardResponse, err error) {
	ts = new(boardResponse)
	uri := fmt.Sprintf("/arrivals/%s", hq)
	err = hux.doRequest(uri, ts)
	return
}
