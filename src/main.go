package main

import (
	"fmt"
	"monkey_interpreter/src/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s This is monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type the commands")
	repl.Start(os.Stdin, os.Stdout)
}
