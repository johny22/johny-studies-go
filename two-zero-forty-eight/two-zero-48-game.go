package main

import "fmt"

func main() {
	const LIMIT int = 2000
	sum := 1

	for sum < LIMIT {
		sum += sum
	}
	fmt.Println(sum)
}
