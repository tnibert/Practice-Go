package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string, more string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "You deny me your name!?", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    // more will be ignored by Sprintf() if no second %v in string
    message := fmt.Sprintf(randomFormat(), name, more)
    return message, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v.  Welcome!  This is your second parameter - %v",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
