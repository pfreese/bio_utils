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

func main() {
	f, _ := os.Open("../pkg/test_data/test_vals_0.csv")
	r := csv.NewReader(bufio.NewReader(f))
	records, _ := r.ReadAll()

	var intMax int
	for _, record := range records {
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			err = errors.Wrap(err, "Parse failed")
			log.Fatal(err)
		}
		if intVal > intMax {
			intMax = intVal
		}
	}
	fmt.Println("Max val = ", intMax)
}
