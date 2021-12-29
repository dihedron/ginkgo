package main

import (
	"fmt"
	"os"

	"github.com/dihedron/ginkgo/ginkgo"
	"github.com/jessevdk/go-flags"
)

func main() {
	ginkgo := ginkgo.Engine{}

	parser := flags.NewParser(&ginkgo, flags.Default)
	if _, err := parser.Parse(); err != nil {
		fmt.Fprintf(os.Stderr, "parsing error: %v\n", err)
		os.Exit(1)
	}
	err := ginkgo.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "execution error: %v\n", err)
		os.Exit(1)
	}
}
