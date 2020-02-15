//
// gobyexample.com
// atomiccounters.go
//
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64
	var waitgroup sync.WaitGroup

	for i := 0; i < 50; i++ {
		waitgroup.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()

	fmt.Println("ops:", ops)
}
