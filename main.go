package main

import (
	"fmt"
	"os"
	"os/user"
	"play/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the play programming language!\n", usr.Username)
	fmt.Println("Test the Lexer")
	repl.Start(os.Stdin, os.Stdout)
}
