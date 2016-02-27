package main

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/datamanager"
	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/dto"
	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/qutils"
)

const url = "amqp://localhost:5672"

func main() {
	conn, ch := qutils.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		qutils.PersistReadingsQueue, //queue string,
		"",    //consumer string,
		false, //autoAck bool,
		true,  //exclusive bool,
		false, //noLocal bool,
		false, //noWait bool,
		nil)   //args amqp.Table)

	if err != nil {
		log.Fatalln("Failed to get access to messages")
	}

	for msg := range msgs {
		buf := bytes.NewReader(msg.Body)
		dec := gob.NewDecoder(buf)
		sd := &dto.SensorMessage{}
		dec.Decode(sd)

		err := datamanager.SaveReading(sd)

		if err != nil {
			log.Printf("Failed to save reading from sensor %v. Error: %s", sd.Name, err.Error())
		} else {
			msg.Ack(false)
		}
	}
}
