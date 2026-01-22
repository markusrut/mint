package main

import (
	"fmt"
	"mint/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! Welcome to MINT\n", user.Username)
	fmt.Printf("Start typing your commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
