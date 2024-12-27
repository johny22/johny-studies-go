package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(num1, num2 byte) byte {
	return num1 + num2
}

func swapProgrammer(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Naked return
}

func main() {
	var diceFace = rand.Intn(5) // rand were namespaced inside math here. The line 5 were required!
	var day1, day2 byte = 12, 21
	var coder1, coder2 = swapProgrammer("Jones", "Josh") // What a awesome pair programming...
	var x, y = split(64)

	fmt.Println("Split 64 => ", x, y)
	fmt.Println("Hey roll my dice! The number is:", diceFace+1) // Dice has no zero side
	fmt.Printf("Now you have %g problems.\nLook that constant %g...\n", math.Sqrt(49), math.Phi)
	fmt.Printf("%d + %d = %d \u03a9\n", day1, day2, add(day1, day2))
	fmt.Printf("Welcome %s! Say go %s! There we go!\n", coder1, coder2)
}
