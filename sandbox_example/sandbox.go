package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome out of the box!")

	fmt.Println("The time is", time.Now())
	fmt.Println("And tomorrow is", time.Now().AddDate(0, 0, 1))
}
