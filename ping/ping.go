package main

import (
  "github.com/bitly/go-nsq"
  "sync"
  "time"
  "fmt"
  "log"
)


type Essence struct{
  config *Config
  sender *Producer
  receiver *Consumer
}

var(
  kaka = true
)


func (ess Essence) init(){
  ess.config = nsq.NewConfig()
  ess.sender, _ := nsq.NewProducer("nsq:4150", ess.config) 
  ess.receiver, _ := nsq.NewConsumer("bong", "ch", config)
}

func (ess Essence) send_ping(){
  err := ess.sender.Publish("bing", []byte("Ping"))
  if err != nil {
      log.Panic("Could not connect wala nada nothing")
  }
  fmt.Println("Ping is sent")
}


func (ess Essence) receive_pong(){
  if kaka{
    kaka = false
    ess.receiver.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
        fmt.Println(string(message.Body), "is received")
        ess.send_ping()
        kaka = true
        return nil
    }))
    err := ess.receiver.ConnectToNSQD("nsq:4150")
    if err != nil {
        log.Panic("Could not connect")
    }
  }
}




func main() {
  ess := new(Essence)
  ess.init()
  ess.send_ping()
  tickChan := time.NewTicker(time.Millisecond * 5000).C
  for {
    select{
      case <- tickChan:
        ess.receive_pong()
    }
  }
}




