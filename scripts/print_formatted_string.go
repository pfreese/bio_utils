package main

import (
	"bytes"
	"fmt"
)

func returnFormattedStringSingleFloat(f float64) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("inputted val:\t%f", f))
	fmt.Println(buffer.String())
	return buffer.String()
}

func returnFormattedStringMultFloats(flts ...float64) string {
	var buffer bytes.Buffer
	buffer.WriteString("Inputted vals:")

	for _, val := range flts {
		buffer.WriteString(fmt.Sprintf("\t%f", val))
	}
	fmt.Println(buffer.String())
	return buffer.String()
}

func main() {
	// prints: "inputted val:   5.400000"
	returnFormattedStringSingleFloat(5.4)

	// prints: "Inputted vals:  4.300000        6.000000        2.100000"
	returnFormattedStringMultFloats(4.3, 6, 2.1)

	// If multiple floats are in a slice, to apply them to a variadic
	// function, use "..."
	// prints: "Inputted vals:  1.000000        2.000000"
	flts := []float64{1, 2}
	returnFormattedStringMultFloats(flts...)
}
