package main

import "fmt"


func main() {
    // var name = "Marcello"
    var age int32 = 23
    const isCool = true

    //shorthand
    // name := "Marcello"
    size := 3.2  
    // email := "marcellovcs@gmail.com"
    name, email := "Marcello", "marcellovcs@gmail.com"

    fmt.Println(name, age, isCool, size, email)
    fmt.Printf("%T\n", email)
}
