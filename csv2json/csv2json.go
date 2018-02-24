package csv2json

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func readFromCsv(r *csv.Reader, c chan *map[string]string) {

	headers, e := r.Read()

	if e != nil {
		log.Fatalln("Can not read the input file.")
		os.Exit(1)
	}

	for i := 0; i < len(headers); i++ {
		headers[i] = strings.ToLower(headers[i])
	}

	for {
		line, e := r.Read()
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatalln(e)
			os.Exit(1)
		} else {
			record := make(map[string]string)
			for idx, h := range headers {
				record[h] = line[idx]
			}
			c <- &record
		}
	}

	close(c)
}

// Transform csv file to json file
func Transform(r *bufio.Reader, w *bufio.Writer) {

	c := make(chan *map[string]string, 10)
	csvReader := csv.NewReader(r)

	go readFromCsv(csvReader, c)

	jsonList := make([]map[string]string, 0)

	for line := range c {
		jsonList = append(jsonList, *line)
	}
	m, _ := json.MarshalIndent(jsonList, "", "    ")
	w.Write(m)
}
