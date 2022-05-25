package parse_test

import (
	"strings"
	"testing"

	"github.com/danielcooperxyz/port-json/parse"
	"github.com/stretchr/testify/assert"
)

func TestParsePortStream(t *testing.T) {
	inputStream := strings.NewReader(testInput)

	records := make(chan interface{}, 5)
	parse.ParseJson(inputStream, records, parse.ParsePort)

	assert.Equal(t, 3, len(records))

	i := 0
	for rec := range records {
		expected := expectations[i]

		port := rec.(*parse.Port)

		assert.Equal(t, expected.ID, port.ID)
		assert.Equal(t, expected.Name, port.Name)
		assert.Equal(t, expected.City, port.City)
		assert.Equal(t, expected.Country, port.Country)
		assert.Equal(t, expected.Province, port.Province)
		assert.Equal(t, expected.Timezone, port.Timezone)
		assert.Equal(t, expected.Code, port.Code)
		assert.Equal(t, expected.Alias, port.Alias)
		assert.Equal(t, expected.Regions, port.Regions)
		assert.Equal(t, expected.Coordinates, port.Coordinates)
		assert.Equal(t, expected.Unlocs, port.Unlocs)

		i++
	}
}

var testInput = `{
	"AEAJM": {
	  "name": "Ajman",
	  "city": "Ajman",
	  "country": "United Arab Emirates",
	  "alias": [],
	  "regions": [],
	  "coordinates": [
		55.5136433,
		25.4052165
	  ],
	  "province": "Ajman",
	  "timezone": "Asia/Dubai",
	  "unlocs": [
		"AEAJM"
	  ],
	  "code": "52000"
	},
	"AEAUH": {
	  "name": "Abu Dhabi",
	  "coordinates": [
		54.37,
		24.47
	  ],
	  "city": "Abu Dhabi",
	  "province": "Abu Z¸aby [Abu Dhabi]",
	  "country": "United Arab Emirates",
	  "alias": [],
	  "regions": [],
	  "timezone": "Asia/Dubai",
	  "unlocs": [
		"AEAUH"
	  ],
	  "code": "52001"
	},
	"AEDXB": {
	  "name": "Dubai",
	  "coordinates": [
		55.27,
		25.25
	  ],
	  "city": "Dubai",
	  "province": "Dubayy [Dubai]",
	  "country": "United Arab Emirates",
	  "alias": [],
	  "regions": [],
	  "timezone": "Asia/Dubai",
	  "unlocs": [
		"AEDXB"
	  ],
	  "code": "52005"
	}
}`

var expectations = []*parse.Port{
	{
		ID:      "AEAJM",
		Name:    "Ajman",
		City:    "Ajman",
		Country: "United Arab Emirates",
		Alias:   []interface{}{},
		Regions: []interface{}{},
		Coordinates: []float32{
			55.5136433,
			25.4052165,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEAJM",
		},
		Code: "52000",
	},
	{
		ID:      "AEAUH",
		Name:    "Abu Dhabi",
		Alias:   []interface{}{},
		Regions: []interface{}{},
		Coordinates: []float32{
			54.37,
			24.47,
		},
		City:     "Abu Dhabi",
		Province: "Abu Z¸aby [Abu Dhabi]",
		Country:  "United Arab Emirates",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEAUH",
		},
		Code: "52001",
	},
	{
		ID:      "AEDXB",
		Name:    "Dubai",
		Alias:   []interface{}{},
		Regions: []interface{}{},
		Coordinates: []float32{
			55.27,
			25.25,
		},
		City:     "Dubai",
		Province: "Dubayy [Dubai]",
		Country:  "United Arab Emirates",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEDXB",
		},
		Code: "52005",
	},
}
