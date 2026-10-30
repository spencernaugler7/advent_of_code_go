package main

import (
	"example/hello/dayOne"
	"example/hello/dayTwo"
	"flag"
	"fmt"

	"github.com/samber/lo"
	"golang.org/x/tools/go/packages"
)

type Blah struct {
	Solutions map[int]func(int)
}

var FunctionMap Blah = Blah{}

func main() {

	var day = flag.Int("day", 0, "day for the advent of code")
	var part = flag.Int("part", 0, "part of the puzzle for the provided day")
	flag.Parse()

	cfg := &packages.Config{Mode: packages.LoadFiles}
	pkgs, _ := packages.Load(cfg, "example/hello/...")

	if *day != 0 && *part != 0 {
		for _, pkg := range pkgs {
			fmt.Println(pkg.Name)
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

	numbers := lo.Range(40)
	squared := lo.Map(numbers, func(blah, _ int) int { return blah * blah })
	sum := lo.Sum(squared)
	fmt.Printf("%d", sum)
}
