package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string, more string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v.  Welcome!  This is your second parameter - %v", name, more)
    return message
}
