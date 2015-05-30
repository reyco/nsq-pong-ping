package main

import (
  "github.com/bitly/go-nsq"
  "sync"
  "time"
  "fmt"
  "log"
)


var(
  kaka = true
)


func send_pong(){
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("nsq:4150", config)

  err := w.Publish("bong", []byte("Pong"))
  if err != nil {
      log.Panic("Could not connect wala")
  }
  fmt.Println("Pong is sent")
  w.Stop()

}

func receive_ping(){
  kaka = false
  wg := &sync.WaitGroup{}
  wg.Add(1)

  config := nsq.NewConfig()
  q, _ := nsq.NewConsumer("bing", "ch", config)
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      fmt.Println(string(message.Body), "is received")
      // wg.Done()
      send_pong()
      kaka = true
      return nil
  }))
  err := q.ConnectToNSQD("nsq:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
  // wg.Wait()
}



func main() {
  tickChan := time.NewTicker(time.Millisecond * 400).C
  for {
    select{
      case <- tickChan:
        if kaka{
	  receive_ping()
	  time.Sleep(time.Millisecond*5000)
	}
    }
  }
}





