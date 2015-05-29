package main

import (
  "github.com/bitly/go-nsq"
  "sync"
  "time"
  "fmt"
  "log"
)


func send_ping(){
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("nsq:4150", config)

  err := w.Publish("bing", []byte("Ping"))
  if err != nil {
      log.Panic("Could not connect wala")
  }
  w.Stop()

}


func receive_pong(){
  wg := &sync.WaitGroup{}
  wg.Add(1)

  config := nsq.NewConfig()
  q, _ := nsq.NewConsumer("bong", "ch", config)
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      fmt.Println(string(message.Body))
      wg.Done()
      send_ping()
      return nil
  }))
  err := q.ConnectToNSQD("nsq:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
  wg.Wait()
}



func main() {
  send_ping()
  for {
    receive_pong()
    time.Sleep(time.Millisecond*500)
  }
}





