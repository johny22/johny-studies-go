package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("Meet my friend... יֵשׁוּעַ") // Yeshua!
	fmt.Println(quote.Go() + "\n\t\u03a9", '!') // It uses ' for sigle character and " for strings :)
}

// Then... Maranata!