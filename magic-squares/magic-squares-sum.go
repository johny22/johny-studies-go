package main

import (
	"fmt"
	"math"
)

func GetMagicSquareSum(slots int) (magicSum int) {
	sum := 0
	for i := 1; i <= int(slots); i++ {
		sum += i
	}
	magicSum_ := int(float64(sum) / math.Sqrt(float64(slots)))
	return magicSum_
}

func main() {
	for j := 2; j < 6; j++ {
		fmt.Printf("Ordem %d! Soma: %d\n", j, GetMagicSquareSum(int(j*j)))
	}
}
