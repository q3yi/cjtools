package csv2json

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"strings"
	"testing"
)

func TestReadFromCSV(t *testing.T) {
	csvStr := "a,b,c\n1,2,3\n"

	c := make(chan *map[string]string, 10)
	csvReader := csv.NewReader(strings.NewReader(csvStr))

	go readFromCsv(csvReader, c)

	jsonList := make([]map[string]string, 0)

	for line := range c {
		jsonList = append(jsonList, *line)
	}

	if jsonList[0]["a"] != "1" {
		t.Error("For", csvStr, "expected", "1", "got", jsonList[0]["a"])
	}
}

func TestTransform(t *testing.T) {
	csvStr := "a,b,c\n1,2,3\n"
	jsonStr := `[
    {
        "a": "1",
        "b": "2",
        "c": "3"
    }
]`
	r := bufio.NewReader(strings.NewReader(csvStr))
	bb := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(bb)
	Transform(r, w)
	w.Flush()
	resultStr := bb.String()
	if jsonStr != resultStr {
		t.Error("For", csvStr, "expect", jsonStr, "got", resultStr)
	}
}
