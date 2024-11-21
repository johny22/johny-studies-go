package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, Jones! Let's GO!\n\n")
	fmt.Println(quote.Go() + "\n\t\u03a9", '!') // It uses ' for sigle character and " for strings :)
	fmt.Println(quote.Opt() + "\n\n")
	fmt.Println(quote.Glass() + "\t:v\n\n")
}