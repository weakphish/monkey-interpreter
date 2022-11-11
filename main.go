package main

import (
	"fmt"
	"github.com/weakphish/monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
