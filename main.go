package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kfess/go_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Fell free to type in commnads\n")
	repl.Start(os.Stdin, os.Stdout)
}
