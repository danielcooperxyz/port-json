package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/danielcooperxyz/port-json/parse"
)

func main() {
	catchKillSignals()

	var filename string

	flag.StringVar(&filename, "filename", "challenge/ports.json", "The filename to parse for port information.")
	flag.Parse()

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
			fmt.Printf("Previously stored: %s\n", p.ID)
		}

		portMap[p.ID] = p
	}
}

func catchKillSignals() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc

		log.Printf("Exit signal '%v' received, closing...\n", s)
		os.Exit(1)
	}()
}
