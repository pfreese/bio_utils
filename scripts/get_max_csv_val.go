package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type Params struct {
	csvSep string
}

func getMaxFromFirstColumn(File string, csvSep rune) {
	f, _ := os.Open(File)
	r := csv.NewReader(bufio.NewReader(f))
	// Set the 'comma' to be csvSep, likely either '\t' or ','
	r.Comma = csvSep

	records, _ := r.ReadAll()

	var intMax int
	for i, record := range records {
		fmt.Println(record)
		// Convert the string to int
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			err = errors.Wrapf(err, "Failed at %d", i) // just .Wrap if not formatting
			log.Fatal(err)
		}
		// Update if this line contains the greatest value seen so far
		if intVal > intMax {
			intMax = intVal
		}
	}
	fmt.Println("Max val = ", intMax)
}

func getMaxFromCSV(File string) {
	getMaxFromFirstColumn(File, ',')
}

func getMaxFromTSV(File string) {
	getMaxFromFirstColumn(File, '\t')
}

func main() {
	// Test a comma-separated file
	testCSV := "../pkg/test_data/test_vals_0.csv"
	getMaxFromCSV(testCSV)

	// Test a tab-separated file
	testTSV := "../pkg/test_data/test_vals_0.tsv"
	getMaxFromTSV(testTSV)
}
