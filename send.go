package main

import (
  "fmt"
  "log"
  "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {                         //Helper function to check the return value for each amqp call

  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}

func main() {

  conn, err := amqp.Dial("amqp://test:test@docker-vm:5672/test") //Start connection with amqp.Dial()
  failOnError(err, "Failed to connect to RabbitMQ")           //Check if an error was returned
  defer conn.Close()                                                //Close connection and wait for response

  ch, err := conn.Channel()                                         //Create a channel
  failOnError(err, "Failed to open a channel")                //Check the return of create channel request
  defer ch.Close()                                                  //Close channel and wait for response

  q, err := ch.QueueDeclare(                                        //Declare a queue that a message can be send to
    "hello", // name
    false,   // durable
    false,   // delete when unused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
  )
  failOnError(err, "Failed to declare a queue")               //Check if queue was declared successfully

  body := "hello"                                                   //Create message body
  err = ch.Publish(                                                 //Publish the message
    "",     // exchange
    q.Name, // routing key
    false,  // mandatory
    false,  // immediate
    amqp.Publishing{
      ContentType: "text/plain",
      Body:        []byte(body),                                    //Pass message body as a byte array
    })
  failOnError(err, "Failed to publish a message")

}

