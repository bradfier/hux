package hux

import (
	"fmt"
)

type huxQuery string

// HuxQueryBuilder constructs a formatted URI query string for use with
// the Huxley API proxy, using a builder pattern to set the required
// and optional parameters. For instance, to query all trains departing
// London Kings Cross (KGX) bound for Cambridge (CBG), construct a query
// as follows:
//
//	hqb := new(HuxQueryBuilder)
//	hq := hqb.QueryStation("KGX").FilterTo().FilterStation("CBG").Build()
//
// A QueryStation() must always be provided as the base of the query, and
// note that the direction of travel must be specified with FilterFrom()
// or FilterTo() if an additional station discriminator is provided.
//
// If you prefer, full station names may be used in place of three character
// CRS codes:
//
//	hqb := new(HuxQueryBuilder)
//	hq := hqb.QueryStation("Cambridge").FilterFrom().FilterStation("Kings Cross").Build()
//
type HuxQueryBuilder struct {
	queryStation    string
	filterDirection string
	filterStation   string
	numRows         int
}

// QueryStation sets the base station for the constructed query, must always be provided.
func (qb *HuxQueryBuilder) QueryStation(station string) *HuxQueryBuilder {
	qb.queryStation = station
	return qb
}

// FilterTo sets the direction of travel to select trains towards the FilterStation discriminator.
func (qb *HuxQueryBuilder) FilterTo() *HuxQueryBuilder {
	qb.filterDirection = "to"
	return qb
}

// FilterFrom sets the direction of travel to select trains from the FilterStation discriminator.
func (qb *HuxQueryBuilder) FilterFrom() *HuxQueryBuilder {
	qb.filterDirection = "from"
	return qb
}

// FilterStation sets the query to select only trains To or From that station.
func (qb *HuxQueryBuilder) FilterStation(station string) *HuxQueryBuilder {
	qb.filterStation = station
	return qb
}

// NumRows limits the number of result rows returned by the API.
func (qb *HuxQueryBuilder) NumRows(rows int) *HuxQueryBuilder {
	qb.numRows = rows
	return qb
}

// Build finalizes the HuxQueryBuilder into a formatted huxQuery query string usable by hux Get methods.
func (qb *HuxQueryBuilder) Build() huxQuery {
	if qb.queryStation == "" {
		panic("Query station must be set")
	}

	str := "/" + qb.queryStation

	// Optional filter parameters
	if qb.filterDirection != "" && qb.filterStation != "" {
		str = str + "/" + string(qb.filterDirection) + "/" + qb.filterStation
	}

	if qb.numRows != 0 {
		str = fmt.Sprintf("%s/%d", str, qb.numRows)
	}

	return huxQuery(str)
}
