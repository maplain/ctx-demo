package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/maplain/ctx-demo/pkg/printer"
)

// +build ignore

func main() {
	var lock sync.Mutex
	var a int
	printer, closePrinter := printer.GetLoadBalancedPrinter(5)
	go func() {
		printer(fmt.Sprintf("current time %s", time.Now()))
		time.Sleep(time.Second * 1)
		lock.Lock()
		a = 1
		lock.Unlock()
		printer(fmt.Sprintf("work done %s", time.Now()))
	}()

	for {
		lock.Lock()
		if a == 1 {
			printer(fmt.Sprintf("worker returned %s", time.Now()))
			break
		}
		lock.Unlock()
		time.Sleep(time.Millisecond * 100)
		printer(fmt.Sprintf("check at %s", time.Now()))
	}
	closePrinter()
}
