package chans

import (
  "fmt"
  "time"
)

func SimpleGreeting(msg string, c chan string) {
  for i := 0; ; i++ {
    c <- fmt.Sprintf("Hi! %s %v", msg, i)
    time.Sleep(time.Duration(1 * time.Second))
  }
}
