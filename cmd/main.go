package main

import (
	"fmt"
	"os"

	"github.com/wowterrorism/passmaster/generate"
)

func main() {
	args := os.Args[1:]

	if !(len(args) == 0) {
		fmt.Println("No args provided")
	}
	g := generate.Generator{
		Alpha:     true,
		Numeric:   true,
		Symb:      true,
		UpperOnly: false,
		LowerOnly: false,
	}

	fmt.Println(g.Generate())
	fmt.Println(g.Generate())
	fmt.Println(g.Generate())
}
