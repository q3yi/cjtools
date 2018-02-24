package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	csv2json "github.com/tsingyi/cjtools/csv2json"
	json2csv "github.com/tsingyi/cjtools/json2csv"
)

func main() {
	transform := flag.String("type", "", "csv2json: c2j or json2csv: j2c")
	outputFile := flag.String("output", "", "output filepath")
	flag.Parse()

	var input *bufio.Reader

	if flag.NArg() == 0 {
		input = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatalf("Can not open input file: %s", flag.Arg(0))
			os.Exit(1)
		}
		input = bufio.NewReader(f)
	}

	var output *bufio.Writer

	if *outputFile != "" {
		f, err := os.OpenFile(*outputFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatalf("Can not open ouput file: %s", *outputFile)
			os.Exit(1)
		}
		output = bufio.NewWriter(f)
	} else {
		output = bufio.NewWriter(os.Stdout)
	}

	switch *transform {
	case "c2j":
		csv2json.Transform(input, output)
	case "j2c":
		json2csv.Transform(input, output)
	default:
		log.Fatalf("Unknown transform type: %s", *transform)
		os.Exit(1)
	}

}
