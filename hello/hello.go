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

    // A slice of names.
    names := []string{"Isis", "Ra", "Darrin"}

    // greeting from my module
    messagemap, err := greetings.Hellos(names)
    if err != nil {
       log.Fatal(err)
    }

    for key, val := range messagemap {
        fmt.Printf("Key: %d, Value: %s\n", key, val)
    }
}
