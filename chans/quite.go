package chans

import (
	"fmt"
	"time"
)

func QuiteGreeting(msg string) (<-chan string, <-chan bool) { // return receive only channel
	c := make(chan string)
	done := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Hi! %s %v", msg, i)
			time.Sleep(time.Duration(4 * time.Second))
			// done after the third greeting
			if i == 3 {
				done <- true
			}
		}
	}()
	return c, done
}
