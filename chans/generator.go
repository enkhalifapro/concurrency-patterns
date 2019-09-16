package chans

import (
  "fmt"
  "time"
)

func GeneratorGreeting(msg string) <-chan string { // return receive only channel
  c := make(chan string)
  go func() {
    for i := 0; ; i++ {
      c <- fmt.Sprintf("Hi! %s %v", msg, i)
      time.Sleep(time.Duration(4 * time.Second))
    }
  }()
  return c
}
