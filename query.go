package hux

import (
	"fmt"
)

type huxQuery string

type QueryBuilder struct {
	queryStation    string
	filterDirection string
	filterStation   string
	numRows         int
}

func (qb *QueryBuilder) QueryStation(station string) *QueryBuilder {
	qb.queryStation = station
	return qb
}

func (qb *QueryBuilder) FilterTo() *QueryBuilder {
	qb.filterDirection = "to"
	return qb
}

func (qb *QueryBuilder) FilterFrom() *QueryBuilder {
	qb.filterDirection = "from"
	return qb
}

func (qb *QueryBuilder) FilterStation(station string) *QueryBuilder {
	qb.filterStation = station
	return qb
}

func (qb *QueryBuilder) NumRows(rows int) *QueryBuilder {
	qb.numRows = rows
	return qb
}

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
