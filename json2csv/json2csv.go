package json2csv

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// Transform json file to csv file
func Transform(r *bufio.Reader, w *bufio.Writer) {
	var jsonList []map[string]interface{}

	rawData, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln("Can not read the input file.")
		os.Exit(1)
	}
	if len(rawData) == 0 {
		log.Fatalln("empty input.")
		os.Exit(1)
	}

	err = json.Unmarshal(rawData, &jsonList)
	if err != nil {
		log.Fatalln("Can not parse the json file")
		os.Exit(1)
	}

	headers := make([]string, 0, len(jsonList[0]))

	for k := range jsonList[0] {
		headers = append(headers, k)
	}

	sort.Strings(headers)

	csvWriter := csv.NewWriter(w)

	csvWriter.Write(headers)

	record := make([]string, len(headers))
	for _, recordMap := range jsonList {
		for i, k := range headers {
			s, _ := recordMap[k].(string)
			record[i] = s
		}
		csvWriter.Write(record)
	}
}
