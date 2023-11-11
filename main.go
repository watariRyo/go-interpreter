package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/watariRyo/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is the Monky Programming Language\n", user.Username)
	fmt.Printf("Feel free to type in the program / command\n")
	repl.Start(os.Stdin, os.Stdout)
}
