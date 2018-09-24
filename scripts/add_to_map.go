package main

import "fmt"

func addToMap() {
	m := make(map[int] int)
	for i := 1; i < 10; i++ {
		fmt.Println(i%2)
		m[i%2] += 1
	}
	fmt.Println(m)
}

func main() {
	addToMap()
}
