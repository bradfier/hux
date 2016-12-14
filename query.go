package hux

import (
	"fmt"
)

type huxQuery string

type HuxQueryBuilder struct {
	queryStation    string
	filterDirection string
	filterStation   string
	numRows         int
}

func (qb *HuxQueryBuilder) QueryStation(station string) *HuxQueryBuilder {
	qb.queryStation = station
	return qb
}

func (qb *HuxQueryBuilder) FilterTo() *HuxQueryBuilder {
	qb.filterDirection = "to"
	return qb
}

func (qb *HuxQueryBuilder) FilterFrom() *HuxQueryBuilder {
	qb.filterDirection = "from"
	return qb
}

func (qb *HuxQueryBuilder) FilterStation(station string) *HuxQueryBuilder {
	qb.filterStation = station
	return qb
}

func (qb *HuxQueryBuilder) NumRows(rows int) *HuxQueryBuilder {
	qb.numRows = rows
	return qb
}

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
