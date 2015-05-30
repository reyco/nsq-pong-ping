package main

import (
  "github.com/bitly/go-nsq"
  "time"
  "fmt"
  "log"
)


type Essence struct{
  config *nsq.Config
  sender *nsq.Producer
  receiver *nsq.Consumer
}


func (ess *Essence) init(){
  ess.config = nsq.NewConfig()
  ess.sender, _ = nsq.NewProducer("nsq:4150", ess.config) 
  ess.receiver, _ = nsq.NewConsumer("bong", "ch", ess.config)
}

func (ess *Essence) send_ping(){
  err := ess.sender.Publish("bing", []byte("Ping"))
  if err != nil {
      log.Panic("Could not connect wala nada nothing")
  }
  fmt.Println("Ping is sent")
}


func (ess *Essence) receive_pong(){
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


var(
  kaka = true
  ess Essence
)



func main() {
  ess.init()
  ess.send_ping()
  // ess.receive_pong()
  for {
    time.Sleep(time.Millisecond * 5000)
  }
}




