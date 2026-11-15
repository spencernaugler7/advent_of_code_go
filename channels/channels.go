package channels

import (
	"fmt"
	"io"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func Register(registry map[string]func(io.Writer)) map[string]func(io.Writer) {
	registry["channels"] = Run
	return registry
}

func Run(writer io.Writer) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Fprintln(writer, <-pongs)
}
