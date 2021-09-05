package main

import (
	"fmt"
	"ironlotuscomputer.com.au/greetings"
	"log"
	"math"
	"rsc.io/quote"
	"time"
)

/*
This has just become a very long series of experiments
*/
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

	// interesting
	num := 3
	for {
		if num < 1 {
			break
		} else {
			num--
		}
	}
	fmt.Printf("the final number is: %d\n", num)

	// also interesting
	if x := 6; x < 5 {
		fmt.Printf("pre if statement set x to %d\n", x)
	} else {
		fmt.Printf("pre if statement variable declaration available in else: %d\n", x)
	}

	// square root algorithm
	fmt.Println(Sqrt(9))

	// switch with no condition is like switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// test defer
	defer_helloworld()
}

func forloop_sum(iters int) (int, error) {
	sum := 0
	for i := 0; i < iters; i++ {
		sum += i
	}
	return sum, nil
}

// reduce a slice
func reduce_sum(myslice []int) int {
	if len(myslice) == 1 {
		return myslice[0]
	}
	return myslice[len(myslice)-1] + reduce_sum(myslice[:len(myslice)-1])
}

func Sqrt(x float64) float64 {
	guess := 1.0
	precision := math.Pow(10, -10)

	fmt.Println("Computing square root")
	for math.Abs(guess*guess-x) > precision {
		guess -= (guess*guess - x) / (2 * guess)
		fmt.Printf("Guess: %.10f\n", guess)
	}
	return guess
}

func defer_helloworld() {
	// world will print when function returns
	defer fmt.Println("world")
	// apparently defer executes in reverse order
	defer fmt.Println("When is this one executed?")
	fmt.Println("hello")
}
