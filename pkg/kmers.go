package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// Return a buffer of runes, e.g.,: [67 84 84]; note
// randKmergenerated := randomKmer(3)
// fmt.Println(randKmergenerated)         >> [67 84 84]
// fmt.Println(string(randKmergenerated)) >> CTT
// Note that the random seed should be set before this is called:
// rand.Seed(time.Now().UTC().UnixNano())
func randomKmer(k int) []byte {
	var buffer bytes.Buffer
	for i := 0; i < k; i++ {
		switch rand.Intn(4) {
		case 0:
			buffer.Write([]byte{'A'})
		case 1:
			buffer.Write([]byte{'C'})
		case 2:
			buffer.Write([]byte{'G'})
		case 3:
			buffer.Write([]byte{'T'})
		}
	}
	return buffer.Bytes()
}

func allKmersRecursive(k int, existingKmers []string) []string {
	if k == 0 {
		return existingKmers
	}
	nts := []string{"A", "C", "G", "T"}
	var extendedKmers []string
	for _, existingKmer := range existingKmers {
		for _, nt := range nts {
			extendedKmers = append(extendedKmers, existingKmer + nt)
		}
	}
	return allKmersRecursive(k - 1, extendedKmers)
}

// Returns a slice containing all kmers in alphabetical order
func allKmers(k int) []string {
	return allKmersRecursive(k, []string{""})
}

func main() {
	// Set random seed so we get a new answer each time - note that
	// it's much faster here than in the randomKmer function
	rand.Seed(time.Now().UTC().UnixNano())
	randKmergenerated := randomKmer(3)
	fmt.Println(randKmergenerated)
	fmt.Println(string(randKmergenerated))

	//fmt.Println(allKmersRecursive(0, []string{"g", "h", "i"}))
	//fmt.Println(allKmersRecursive(1, []string{"g", "h", "i"}))
	fmt.Println(allKmers(3))
}
