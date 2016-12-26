package hux

import (
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	hqb := new(QueryBuilder)
	hq := hqb.QueryStation("KGX").FilterFrom().FilterStation("CBG").Build()

	if hq != "/KGX/from/CBG" {
		t.Error()
	}

	hq = hqb.QueryStation("Cambridge").FilterFrom().FilterStation("Kings Cross").
		NumRows(15).Build()

	if hq != "/Cambridge/from/Kings Cross/15" {
		t.Error()
	}
}
