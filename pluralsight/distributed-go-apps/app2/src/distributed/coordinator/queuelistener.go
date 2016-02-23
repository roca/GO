package coordinator

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/dto"
	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/qutils"
	"github.com/streadway/amqp"
)

const url = "amqp://localhost:5672"

type QueueListener struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	sources map[string]<-chan amqp.Delivery
	ea      *EventAggregator
}

func NewQueueListener() *QueueListener {
	ql := QueueListener{
		sources: make(map[string]<-chan amqp.Delivery),
		ea:      NewEventAggregator(),
	}

	ql.conn, ql.ch = qutils.GetChannel(url)
	ql.sources = make(map[string]<-chan amqp.Delivery)

	return &ql
}

func (ql *QueueListener) DiscoverSensors() {
	ql.ch.ExchangeDeclare(
		qutils.SensorDiscoveryExchange, //name string,
		"fanout",                       //kind string,
		false,                          //durable bool,
		false,                          //autoDelete bool,
		false,                          //internal bool,
		false,                          //noWait bool,
		nil)

	ql.ch.Publish(
		qutils.SensorDiscoveryExchange, //exchange string,
		"",                //key string,
		false,             //mandatory bool,
		false,             //immediate bool,
		amqp.Publishing{}) //args amqp.Table)
}

func (ql *QueueListener) ListenForNewSource() {
	q := qutils.GetQueue("", ql.ch)
	ql.ch.QueueBind(
		q.Name,       //Name string
		"",           //Key string
		"amq.fanout", //exchange string
		false,        //noWait bool
		nil)          //args amqp.Table

	msgs, _ := ql.ch.Consume(
		q.Name, //queue string,
		"",     //consumer string,
		true,   //autoAck bool,
		false,  //exclusive bool,
		false,  //noLocal bool,
		false,  //noWait bool,
		nil)    //args amqp.Table)

	ql.DiscoverSensors()

	fmt.Println("listening for new sources")
	for msg := range msgs {
		fmt.Println("new source discovered")
		sourceChan, _ := ql.ch.Consume(
			string(msg.Body), //queue string,
			"",               //consumer string,
			true,             //autoAck bool,
			false,            //exclusive bool,
			false,            // noLocal bool,
			false,            // noWait bool,
			nil)              // args amqp.Table)

		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan

			go ql.AddListener(sourceChan)
		}
	}
}

func (ql *QueueListener) AddListener(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(dto.SensorMessage)
		d.Decode(sd)

		ed := EventData{
			Name:      sd.Name,
			Value:     sd.Value,
			Timestamp: sd.Timestamp,
		}

		ql.ea.PublishEvent("MessageRecieved_"+msg.RoutingKey, ed)

		fmt.Printf("Recieved message: %v\n", sd)
	}
}
