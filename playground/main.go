// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `select` builtin on the channel to await the
	// values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	wg:=&sync.WaitGroup{}


	run:= func(wg *sync.WaitGroup, num int) {
		for {
			select {
			case d:=<-done:
				fmt.Println(d,":wg done")
				wg.Done()
				return
			case t := <-ticker.C:
				fmt.Println(num,":Tick at", t)
			}
		}
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go run(wg,i+1)
	}

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	close(done)
	//done <- true
	fmt.Println("Ticker stopped")
	wg.Wait()
	os.Exit(0)
}
