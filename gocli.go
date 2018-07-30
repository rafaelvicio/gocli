package main

import (
	"flag"
	"os"

	"github.com/rafaelvicio/goclit/generate"
)

func main() {
	typeFlag := flag.String("type", "", "Type of Bilhete.")
	outputFlag := flag.String("output", "", "Output name, default: ./output.xml")

	flag.Parse()

	if *typeFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	generate.Generate(*typeFlag, *outputFlag)

}
