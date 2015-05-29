package main

import (
  "log"
  "github.com/bitly/go-nsq"
  "time"
)

func main() {
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("nsq:4150", config)

  err := w.Publish("kakang", []byte("testa" + time.Now().String() ))
  if err != nil {
      log.Panic("Could not connect wala")
  }
  
  println("Transmitted Ping");
  
  w.Stop()
}
