package main

import (
    "fmt"
    "rsc.io/quote"
    "ironlotuscomputer.com.au/greetings"
)

func main() {
    // quote from quote library
    fmt.Println(quote.Go())

    // greeting from my library
    message := greetings.Hello("Tim", "Boom shacka lacka")
    fmt.Println(message)
}
