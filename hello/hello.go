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
    res, err := forloop_sum(2)
    fmt.Printf("sum: %d\n", res)

    s := []int{1, 2, 3, 4}
    res = reduce_sum(s)
    fmt.Printf("reduced: %d\n", res)
}

func forloop_sum(iters int) (int, error) {
    sum := 0
    for i:= 0; i < iters; i++ {
        sum += i
    }
    return sum, nil
}

// reduce a slice
func reduce_sum(myslice []int) (int) {
    if len(myslice) == 1 {
        return myslice[0]
    }
    return myslice[len(myslice)-1] + reduce_sum(myslice[:len(myslice)-1])
}