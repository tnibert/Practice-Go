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

    fmt.Println("Testing for loop")
    res, err := forloop_sum(10)
    fmt.Printf("sum: %d\n", res)
}

func forloop_sum(iters int) (int, error) {
    sum := 0
    for i:= 0; i < iters; i++ {
        sum += i
    }
    return sum, nil
}

// todo - recursion
//fun reduce_sum(iters int) (int) {
//    if basecase
//       return
//    return recurse
//}