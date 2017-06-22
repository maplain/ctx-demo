package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/maplain/ctx-demo/pkg/printer"
)

// +build ignore

func main() {
	var lock sync.Locker
	var a int
	lock = &sync.Mutex{}
	cond := sync.NewCond(lock)

	printer, closePrinter := printer.GetPrinter()
	go func() {
		printer(fmt.Sprintf("current time %s", time.Now()))
		time.Sleep(time.Second * 1)
		cond.L.Lock()
		a = 1
		cond.L.Unlock()
		printer(fmt.Sprintf("work done %s", time.Now()))
		cond.Signal()
	}()

	cond.L.Lock()
	for a != 1 {
		printer(fmt.Sprintf("check at %s", time.Now()))
		cond.Wait()
	}
	cond.L.Unlock()
	printer(fmt.Sprintf("worker returned %s", time.Now()))
	closePrinter()
}
