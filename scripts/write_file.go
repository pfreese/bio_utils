package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	outFile := "../pkg/test_outs/test_write_file.txt"
	file, err := os.Create(outFile)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "Hello, readers of bio_utils!")
}
