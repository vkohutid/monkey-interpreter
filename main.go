package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming langauge!\n", user.Username)
	fmt.Print("Feel free t type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
