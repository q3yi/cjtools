# Convert between json and csv file

Convert flaten list json and to csv, or convert csv file to json list.

Only support flaten json object list, like:

```json
[
    {
        "a": 1,
        "b": 2,
        "c": 3
    },
    {
        "a": 4,
        "b": 5,
        "c": 6
    }
]
```

## Usage

* If no output given `stdout` will be used.
* If no input file given `stdin` will be used.

```shell
$> cjtools -type c2j -output output.json input.csv
```

```text
Usage of cjtools:
  -output string
        output filepath
  -type string
        csv2json: c2j or json2csv: j2c
```