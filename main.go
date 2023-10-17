package main

import (
	"flag"
	"fmt"
	"github.com/alissa-tung/glean/glean"
)

func main() {
	glean.InitFlags()

	switch *glean.Command {
	case "elan":
		glean.InstallElan()

	case "lean":
		glean.InstallLean()

	default:
		fmt.Println("unknown command")
		flag.Usage()
	}
}