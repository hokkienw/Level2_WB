package main

import (
	"fmt"
	"time"
)


func or(channels ...<-chan interface{}) <-chan interface{} {
	orDone := make(chan interface{})

	go func() {
		defer close(orDone)
		var cases []<-chan interface{}
		for _, ch := range channels {
			cases = append(cases, ch)
		}

		select {
		case <-cases[0]:
			return
		case <-cases[1]:
			return
		case <-cases[2]:
			return
		case <-cases[3]:
			return
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
