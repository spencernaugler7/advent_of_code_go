package main

import (
	"example/hello/channels"
	"example/hello/fragments"
	"flag"
	"fmt"
)

func main() {
	var task = flag.String("task", "", "enter module name to run")
	flag.Parse()

	if *task == "" {
		return
	}

	registrar := make(map[string]func())
	channels.Register(registrar)
	fragments.Register(registrar)

	run, ok := registrar[*task]

	if !ok {
		panic(fmt.Errorf("Registrar doesn't contain key %s did you register this?", *task))
	}

	run()
}
