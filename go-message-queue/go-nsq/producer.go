package main

// just need to tell which topic and provide data

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	//for i:= 0; i < 1000; i++ {
		//err := producer.Publish("test-nsq", []byte("test" + strconv.Itoa(i)))
		err = producer.Publish("clicks", []byte("haloo"))
		if err != nil {
			log.Fatal(err)
		}
	//}
	producer.Stop()
}