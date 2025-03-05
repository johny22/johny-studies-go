package main

import "fmt"


func Add(num1, num2 int) (value int) {
	return num1 + num2
}

func Substract(num1, num2 int) (value int) {
	return num1 - num2
}

func main() {
	const (
		PLUS int = 1
		MINUS int = 2
		EXIT   int = 3
	)
	var op int
	var menu string = "\n* MENU *\n" +
		"1 - Plus\n" +
		"2 - Minus\n" +
		"3 - Exit\n" +
		"Option: "

	for op != EXIT {
		fmt.Print(menu)
		fmt.Scanln(&op)

		switch op {
			case PLUS:
			{
				var num1, num2 int
				fmt.Print("Enter num 1: ")
				fmt.Scanln(&num1)
				fmt.Print("Enter num 2: ")
				fmt.Scanln(&num2)
				fmt.Printf("The sum of %d + %d = %d.\n", num1, num2, Add(num1, num2))
			}
			case MINUS:
			{
				var num1, num2 int
				fmt.Print("Enter num 1: ")
				fmt.Scanln(&num1)
				fmt.Print("Enter num 2: ")
				fmt.Scanln(&num2)
				fmt.Printf("The subtraction of %d - %d = %d.\n", num1, num2, Substract(num1, num2))
			}
			case EXIT:
				fmt.Println("Exiting now...")
			default:
				fmt.Println("This option doesn't exist!")
		}
	}
}