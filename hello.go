package main

import (
	"example/hello/channels"
	"example/hello/dayOne"
	"example/hello/dayTwo"
	"example/hello/errorhandling"
	"example/hello/fragments"
	"example/hello/lotest"
	"example/hello/multipleinterfaces"
	"flag"
	"fmt"
	"os"
)

func main() {
	var task = flag.String("task", "", "enter module name to run")
	flag.Parse()

	if *task == "" {
		fmt.Fprintln(os.Stderr, "No module provided. Provide flag \"-task\" to pick something to run")
		os.Exit(1)
	}

	registrar := make(map[string]func())
	channels.Register(registrar)
	fragments.Register(registrar)
	dayOne.Register(registrar)
	dayTwo.Register(registrar)
	errorhandling.Register(registrar)
	multipleinterfaces.Register(registrar)
	lotest.Register(registrar)

	run, ok := registrar[*task]

	if !ok {
		fmt.Fprintf(os.Stderr, "Registrar doesn't contain key \"%s\" did you register this?\n", *task)
		os.Exit(1)
	}

	run()
}
