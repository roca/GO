package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const SensorDiscoveryExchange = "SensorDiscovery"
const SensorListQueue = "SensorList"
const PersistReadingsQueue = "PersistReadings"

const WebappSourceExchange = "WebappSources"
const WebappReadingsExchange = "WebappReading"
const WebappDiscoveryQueue = "WebappDiscovery"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,       //name string,
		false,      //durable bool,
		autoDelete, //autoDelete bool,
		false,      //exclusive bool,
		false,      //noWait bool,
		nil)        //args amqp.Table)

	failOnError(err, "Failed to declare queue")

	return &q

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
