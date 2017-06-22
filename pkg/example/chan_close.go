package main

import (
	"fmt"
	"time"

	"github.com/maplain/ctx-demo/pkg/printer"
)

// +build ignore

func main() {
	done := make(chan struct{})

	printer, closePrinter := printer.GetPrinter()
	go func() {
		printer(fmt.Sprintf("current time %s", time.Now()))
		time.Sleep(time.Second * 1)
		printer(fmt.Sprintf("work done %s", time.Now()))
		close(done)
	}()

	<-done
	printer(fmt.Sprintf("worker returned %s", time.Now()))
	closePrinter()
}
