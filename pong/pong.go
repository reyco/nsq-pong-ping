package main

import (
  "github.com/bitly/go-nsq"
  "sync"
  "time"
  "fmt"
  "log"
)


func send_pong(){
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("nsq:4150", config)

  err := w.Publish("bong", []byte("Pong"))
  if err != nil {
      log.Panic("Could not connect wala")
  }
  w.Stop()

}


func receive_ping(){
  wg := &sync.WaitGroup{}
  wg.Add(1)

  config := nsq.NewConfig()
  q, _ := nsq.NewConsumer("bing", "ch", config)
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      fmt.Println(string(message.Body))
      wg.Done()
      time.Sleep(time.Millisecond*500)
      send_pong()
      return nil
  }))
  err := q.ConnectToNSQD("nsq:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
  wg.Wait()
}



func main() {
  for {
    receive_ping()
    time.Sleep(time.Millisecond*500)
  }
}





