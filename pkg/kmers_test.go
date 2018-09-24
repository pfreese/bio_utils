package main

import (
	"fmt"
	"testing"
)

func TestRandomKmer(t *testing.T) {
	KmerToCount := make(map[string] int)
	for i := 1; i < 100; i++ {
		KmerToCount[string(randomKmer(2))] += 1
	}
	fmt.Println(KmerToCount)
	// Make sure that all kmers are present
}
