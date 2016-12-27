package hux

import (
	"fmt"
)

type huxQuery string

// QueryBuilder constructs a filter string, allowing the caller to specify optional
// filters for the departure board data.
type QueryBuilder struct {
	queryStation    string
	filterDirection string
	filterStation   string
	numRows         int
}

// QueryStation sets the station for which the arrivals and departures are to be retrieved.
func (qb *QueryBuilder) QueryStation(station string) *QueryBuilder {
	qb.queryStation = station
	return qb
}

// FilterTo sets the query to destination mode.
func (qb *QueryBuilder) FilterTo() *QueryBuilder {
	qb.filterDirection = "to"
	return qb
}

// FilterFrom sets the query to origin mode.
func (qb *QueryBuilder) FilterFrom() *QueryBuilder {
	qb.filterDirection = "from"
	return qb
}

// FilterStation sets the origin or destination filter as selected by FilterFrom or FilterTo
func (qb *QueryBuilder) FilterStation(station string) *QueryBuilder {
	qb.filterStation = station
	return qb
}

// NumRows limits the number of train services listed in the query response.
func (qb *QueryBuilder) NumRows(rows int) *QueryBuilder {
	qb.numRows = rows
	return qb
}

// Build finalises the parameters in the QueryBuilder into a formatted huxQuery string.
func (qb *QueryBuilder) Build() huxQuery {
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
