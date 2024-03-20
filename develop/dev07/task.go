package main

import (
	"fmt"
	"time"
	"sync"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	orDone := make(chan interface{})

	switch len(channels) {
	case 0:
		out := make(chan interface{})
		close(out)
		return out
	case 1:
		return channels[0]
	}

	go func() {
		var once sync.Once
		for _, c := range channels {
			go func(ch <-chan interface{}) {
				select {
				case _, ok := <-ch:
					if !ok {
						once.Do(func() {
							close(orDone)
						})
					}
				}
			}(c)
		}
	}()
	return orDone
}


func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done %v\n", time.Since(start))
}
