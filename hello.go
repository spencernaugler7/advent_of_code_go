package main

import (
	"fmt"
	"io"
	"os"

	"example/hello/channels"
	"example/hello/dayOne"
	"example/hello/dayTwo"
	"example/hello/defaulttypevalues"
	ecstest "example/hello/ecsTest"
	"example/hello/errorhandling"
	"example/hello/fragments"
	"example/hello/lotest"
	"example/hello/multipleinterfaces"
	"example/hello/ui"
)

func main() {
	registrar := make(map[string]func(io.Writer))
	channels.Register(registrar)
	fragments.Register(registrar)
	dayOne.Register(registrar)
	dayTwo.Register(registrar)
	errorhandling.Register(registrar)
	multipleinterfaces.Register(registrar)
	lotest.Register(registrar)
	defaulttypevalues.Register(registrar)
	ecstest.Register(registrar)

	p := ui.New(&registrar)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
