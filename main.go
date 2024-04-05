package main

import (
	"fmt"
	"os"
	"os/user"

	"cool/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wassup %s? This is the cooles programming language - Cool!\n",
		user.Username)
	fmt.Printf("You may start typin\n")
	repl.Start(os.Stdin, os.Stdout)
}
