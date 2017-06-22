package printer

import (
	"fmt"
	"sync"
)

func GetLoadBalancedPrinter(num int) (printer func(string), closePrinter func()) {
	display := make(chan string)
	var wg sync.WaitGroup

	printer = func(s string) {
		display <- s
	}
	closePrinter = func() {
		close(display)
		wg.Wait()
	}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case s, ok := <-display:
					if !ok {
						return
					}
					fmt.Println(s, " from printer ", id)
				}
			}
		}(i)
	}
	return
}
