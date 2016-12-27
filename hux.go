// Package hux provides Go bindings for the Huxley National Rail API proxy.
package hux

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Hux holds the connection details for the Huxley proxy in use.
type Hux struct {
	baseURI     string
	accessToken string
}

// NewHux returns a Hux with the internal fields populated
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

// GetCRSCodes returns a slice of CRSStationCodes matching the supplied
// filter string. An empty filter string will return all CRS Codes in the National Rail database.
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

// GetAllCRSCodes returns a slice containing all CRS codes in the National Rail database,
// it is functionally equivalent to calling GetCRSCodes with an empty string as the parameter.
func (hux *Hux) GetAllCRSCodes() (stationCodes *[]CRSStationCode, err error) {
	return hux.GetCRSCodes("")
}

// GetDepartures returns a BoardResponse containing the departing trains which match the criteria
// specified in the query.
func (hux *Hux) GetDepartures(hq huxQuery) (ts *BoardResponse, err error) {
	ts = new(BoardResponse)
	uri := fmt.Sprintf("/departures/%s", hq)
	err = hux.doRequest(uri, ts)
	return
}

// GetArrivals returns a BoardResponse containing the arriving trains which match the criteria
// specified in the query.
func (hux *Hux) GetArrivals(hq huxQuery) (ts *BoardResponse, err error) {
	ts = new(BoardResponse)
	uri := fmt.Sprintf("/arrivals/%s", hq)
	err = hux.doRequest(uri, ts)
	return
}

// GetAll returns a BoardResponse containing all the trains which match the criteria
// specified in the query.
func (hux *Hux) GetAll(hq huxQuery) (ts *BoardResponse, err error) {
	ts = new(BoardResponse)
	uri := fmt.Sprintf("/all/%s", hq)
	err = hux.doRequest(uri, ts)
	return
}

// GetNext returns a BoardResponse containing the next train which matches the critera
// specified in the query.
func (hux *Hux) GetNext(hq huxQuery) (ts *BoardResponse, err error) {
	ts = new(BoardResponse)
	uri := fmt.Sprintf("/next/%s", hq)
	err = hux.doRequest(uri, ts)
	return
}

// GetFastest returns a BoardResponse containing the train which will arrive first at
// the destination specified in the query.
func (hux *Hux) GetFastest(hq huxQuery) (ts *BoardResponse, err error) {
	ts = new(BoardResponse)
	uri := fmt.Sprintf("/fastest/%s", hq)
	err = hux.doRequest(uri, ts)
	return
}

