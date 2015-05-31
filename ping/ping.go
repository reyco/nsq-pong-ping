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
  sender_channel string
  receiver_channel string
  sender_message string
}


func (ess *Essence) init(){
  ess.sender_channel = "king"
  ess.receiver_channel = "kong"
  ess.sender_message = "Ping"
  ess.config = nsq.NewConfig()
  ess.sender, _ = nsq.NewProducer("nsq:4150", ess.config) 
  ess.receiver, _ = nsq.NewConsumer(ess.receiver_channel, "ch", ess.config)
}

func (ess *Essence) sending(){
  err := ess.sender.Publish(ess.sender_channel, []byte(ess.sender_message))
  if err != nil {
      log.Panic("Could not connect wala nada nothing")
  }
  fmt.Println(ess.sender_message+" is sent")
}


func (ess *Essence) receiving(){
  ess.receiver.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
    fmt.Println(string(message.Body), "is received")
    time.Sleep(time.Millisecond * 2500)
    ess.sending()
    time.Sleep(time.Millisecond * 2500)
    return nil
  }))
  err := ess.receiver.ConnectToNSQD("nsq:4150")
  if err != nil {
    log.Panic("Could not connect")
  }
}


var(
  ess Essence
)



func main() {
  ess.init()
  ess.sending()
  ess.receiving()
  for {
    time.Sleep(time.Millisecond * 5000)
  }
}




