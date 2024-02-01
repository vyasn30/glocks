package main

import (
	"fmt"
	"os"

	"github.com/vyasn30/glocks/src/glocks"
)

func main() {

	glx := glocks.NewGlocksInterpreter("")

	if len(os.Args) > 2 {
		fmt.Println("Usage: glocks[script]")
	} else if len(os.Args) == 2 {
		glx.RunFile(os.Args[1])
	} else {
		glx.RunPrompt()
	}
}
