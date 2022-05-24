package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "challenge/ports.json", "The filename to parse for port information.")

	log.Printf("Reading %s...\n", filename)

	readJsonFile(filename)
}

type Port struct {
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

func readJsonFile(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(f)

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	for dec.More() {
		// read name
		t, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v\n", t, t)
		p := Port{}
		err = dec.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%#v\n", p)
		return
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}
