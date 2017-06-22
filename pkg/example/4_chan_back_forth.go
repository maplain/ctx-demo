package main

import (
	"fmt"
	"time"

	"github.com/maplain/ctx-demo/pkg/printer"
)

// +build ignore

func main() {
	done := make(chan chan struct{})

	printer, closePrinter := printer.GetPrinter()
	go func() {
		for {
			select {
			case stop := <-done:
				printer(fmt.Sprintf("work done %s", time.Now()))
				stop <- struct{}{}
				return
			default:
				printer(fmt.Sprintf("current time %s", time.Now()))
				time.Sleep(time.Millisecond * 300)
			}
		}
	}()

	time.Sleep(time.Second * 1)
	stop := make(chan struct{})
	done <- stop
	<-stop
	printer(fmt.Sprintf("worker returned %s", time.Now()))
	closePrinter()
}
