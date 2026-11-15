package ui

import "strings"

type WriteCollector struct {
	Output string
}

func (c *WriteCollector) Write(p []byte) (n int, err error) {
	c.Output += strings.Clone(string(p))
	return len(p), nil
}
