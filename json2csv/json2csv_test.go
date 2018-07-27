package json2csv

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestTransform(t *testing.T) {
	istr := `[
		{
			"a": "1",
			"b": "2",
			"c": "3"
		}
	]`
	ecsv := "a,b,c\n1,2,3\n"

	input := bufio.NewReader(strings.NewReader(istr))
	bb := bytes.NewBuffer(make([]byte, 0))
	output := bufio.NewWriter(bb)

	Transform(input, output)

	output.Flush()

	rs := bb.String()

	if rs != ecsv {
		t.Error("For", istr, "expect", ecsv, "got", rs)
	}

}
