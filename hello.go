package main

import (
	"example/hello/dayOne"
	"example/hello/dayTwo"
	"flag"
	"fmt"

	"golang.org/x/tools/go/packages"
)

type Blah struct {
	Solutions map[int]func(int)
}

var FunctionMap Blah = Blah{}

func main() {

	var day = flag.Int("day", 1, "day for the advent of code")
	var part = flag.Int("part", 1, "part of the puzzle for the provided day")
	flag.Parse()

	cfg := &packages.Config{Mode: packages.LoadFiles}
	pkgs, _ := packages.Load(cfg, "example/hello/...")

	if day != nil && part != nil {
		for _, pkg := range pkgs {
			fmt.Println(pkg.Name)
		}
	}

	dayVal := *day
	partVal := *part
	switch {
	case dayVal == 1 && partVal == 1:
		dayOne.Solve(*part)
	case dayVal == 2, partVal == 1:
		dayTwo.Solve(*part)
	}
}
