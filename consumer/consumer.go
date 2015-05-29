package main

import (
  "log"
  "sync"
  "fmt"

  "github.com/bitly/go-nsq"
)

func main() {

  wg := &sync.WaitGroup{}
  wg.Add(1)

  config := nsq.NewConfig()
  q, _ := nsq.NewConsumer("kakang", "ch", config)
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      // log.Printf("Got a message: %v", *message)
      fmt.Println(string(message.Body))
      wg.Done()
      return nil
  }))
  err := q.ConnectToNSQD("nsq:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
  wg.Wait()

}
