package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("Meet my friend... יֵשׁוּעַ") // Yeshua!
	fmt.Println(quote.Go()+"\n\t\u03a9", '!') // It uses ' for sigle character and " for strings :)

	var names = [4]string{"Juana", "Judite", "Jacinta", "Josefa"}
	myslice1 := []int{} //What?
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "stranges", "are", "fast,", "but goodies. Squouiiiirrel!"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

// Then... Maranata!
