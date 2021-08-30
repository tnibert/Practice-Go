package main

import (
    "fmt"
    "log"
    "rsc.io/quote"
    "ironlotuscomputer.com.au/greetings"
)

func main() {
    // quote from quote module
    fmt.Println(quote.Go())

    // greeting from my module
    message, err := greetings.Hello("Test", "Boom shacka lacka")
    if err != nil {
       log.Fatal(err)
    }
    fmt.Println(message)
}
