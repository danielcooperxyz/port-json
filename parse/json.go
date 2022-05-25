package parse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type RecordFunc func(*json.Decoder) interface{}

func ParseJson(input io.Reader, records chan interface{}, parseRecord RecordFunc) {
	dec := json.NewDecoder(input)

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	for dec.More() {
		record := parseRecord(dec)
		records <- record
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	close(records)
}
