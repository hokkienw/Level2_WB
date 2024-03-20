package main

import (
	"testing"
)

func TestOr(t *testing.T) {
	t.Run("AllChannelsClosed", func(t *testing.T) {
		ch1 := make(chan interface{})
		ch2 := make(chan interface{})
		result := or(ch1, ch2)

		go func() {
			close(ch1)
			close(ch2)
		}()

		_, ok := <-result
		if ok {
			t.Error("Expected closed channel, but got a open channel")
		}
	})
}