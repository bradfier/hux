# Hux
[![GoDoc](https://godoc.org/github.com/bradfier/hux?status.svg)](https://godoc.org/github.com/bradfier/hux)
[![Build Status](https://travis-ci.org/bradfier/hux.svg?branch=master)](https://travis-ci.org/bradfier/hux)

Golang bindings for the [Huxley](https://huxley.unop.uk/) JSON proxy by [James Singleton](https://github.com/jpsingleton)

Huxley is a REST proxy for the Darwin National Rail Departure Board SOAP API.

## Usage
First, register for a [Darwin API Key](http://realtime.nationalrail.co.uk/OpenLDBWSRegistration/)

Then, having set up your instance of the Huxley proxy, use Hux to connect from Go:


`````go
import (
	"fmt"
	"github.com/bradfier/hux"
)

func main() {
	accessToken := "abcd1234-ef56-7890-aabb-ccddeeff0011"
	uri := "http://huxley.mydomain.com"

	hux_conn := hux.NewHux(uri, accessToken)
	hq := new(hux.QueryBuilder).QueryStation("KGX").NumRows(10).Build()

	ts, err := hux_conn.GetDepartures(hq)

	if err != nil {
		fmt.Printf("Failed to retrieve departures from Darwin: %s\n", err)
		return
	}

	for _, service := range ts.TrainServices {
		fmt.Printf("Destination: %s  Time: %s  Platform: %s\n",
			service.Destination[0].LocationName, service.ETD, service.Platform)
	}
}
`````

## License
This library is licensed under the BSD 3-Clause license.