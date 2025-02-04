package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

type fiambre int

var (
	ToBe      bool       = 1 == 1
	MaxInt    uint64     = 1<<64 - 1
	Aurum     float64    = math.Phi
	z         complex128 = cmplx.Sqrt(-5 + 12i)
	x, y      rune       = '🐺', '🐍' // Unicode code point rune is an alias for int32
	falseMeat fiambre    = 10
	bread     int        = 11
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Aurum, Aurum)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", falseMeat, falseMeat)
	fmt.Printf("My hotdog costs R$ %v.\n", int(falseMeat)+bread) // OMG that's not even healthy
	fmt.Printf("Type: %T Test =:> %c & %c => %d\n", y, x, y, x+y)
}
