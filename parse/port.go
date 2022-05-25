package parse

import (
	"encoding/json"
	"fmt"
	"log"
)

type Port struct {
	ID string

	Name     string
	City     string
	Country  string
	Province string
	Timezone string
	Code     string

	Alias       []interface{}
	Regions     []interface{}
	Coordinates []float32
	Unlocs      []string
}

func ParsePort(dec *json.Decoder) interface{} {
	// read name
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	p := Port{}
	err = dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	p.ID = fmt.Sprintf("%s", t)
	return &p
}
