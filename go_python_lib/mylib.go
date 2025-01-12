package main

import "C"
import "fmt"

func SayHello(name *C.char) {
    fmt.Println("Hello,", C.GoString(name))
}

func main() {}

