package printer

import "fmt"

func GetPrinter() (printer func(string), closePrinter func()) {
	display := make(chan string)
	done := make(chan struct{})
	printer = func(s string) {
		display <- s
	}
	closePrinter = func() {
		close(display)
		<-done
	}
	go func() {
		for {
			select {
			case s, ok := <-display:
				if !ok {
					close(done)
					return
				}
				fmt.Println(s)
			}
		}
	}()
	return
}
