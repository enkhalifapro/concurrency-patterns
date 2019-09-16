package main

import (
  "concurrency-patterns/chans"
  "fmt"
  "os"
)

func main() {
  fmt.Println("Hello Patterns")
  c := make(chan string)
  go chans.SimpleGreeting("Basic ", c)
  generatorResult := chans.GeneratorGreeting("Generator")
  quiteResult, done := chans.QuiteGreeting("Quite")
  for {
    select {
    case v1 := <-c:
      fmt.Println("from basic: ", v1)
    case v2 := <-generatorResult:
      fmt.Println("from generator: ", v2)
    case v3 := <-quiteResult:
      fmt.Println("from quite: ", v3)
    case <-done:
      fmt.Println("I'm done.....")
      os.Exit(0)
    }
  }
}
