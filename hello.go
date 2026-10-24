package main

import (
	// "flag"
	"fmt"

	"golang.org/x/tools/go/packages"
)

func main() {

	// var day = flag.Int("day", 1, "day for the advent of code")
	// var part = flag.Int("part", 1, "part of the puzzle for the provided day")
	// flag.Parse()

	cfg := &packages.Config{Mode: packages.LoadFiles}
	pkgs, _ := packages.Load(cfg, "example/hello/dayOne")

	for _, pkg := range pkgs {
		for scope := range pkg.Types.Scope().Children() {
			fmt.Println(scope.Names())
		}
	}


	// switch *day {
	// case 1:
		// dayOne.Solve(*part)
	// case 2:
	// 	dayTwo.Solve(*part)
	// }

	// data := map[string]string {
	// 	"day": IntToString(*day),
	// 	"part": IntToString(*part),
	// }
}

func IntToString(num int) string {
	returnVal := ""
	switch num {
	case 1: returnVal = "One"
	case 2: returnVal = "Two"
	case 3: returnVal = "Three"
	case 4: returnVal = "Four"
	case 5: returnVal = "Five"
	}
	return returnVal
}
