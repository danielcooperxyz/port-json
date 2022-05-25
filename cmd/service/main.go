package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danielcooperxyz/port-json/parse"
)

func main() {
	var filename string

	flag.StringVar(&filename, "filename", "challenge/ports.json", "The filename to parse for port information.")
	log.Printf("Reading %s...\n", filename)

	records := make(chan interface{}, 1000)
	portMap := map[string]*parse.Port{}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	go parse.ParseJson(f, records, parse.ParsePort)

	for rec := range records {
		p := rec.(*parse.Port)

		if _, ok := portMap[p.ID]; ok {
			fmt.Printf("Found: %s\n", p.ID)
		}

		portMap[p.ID] = p
	}
}
