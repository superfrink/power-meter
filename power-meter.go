package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/brian-armstrong/gpio"
)

const (
	INPUT_PIN = 27 // GPIO PIN in software (not the hardware pin number)

	LISTEN_HOST = "0.0.0.0" // tcp interface to serve counter on. eg: "0.0.0.0", "127.0.0.1"
	LISTEN_PORT = "9001"    // port to serve counter on
)

// The Counter provides two channels.  The first
// is used to increment a counter.  The second is
// used to read the value.
type Counter struct {
	Increment chan int
	Query     chan chan int
}

func CreateTickCounter() Counter {
	// Creates a new Counter.
	//
	// Sending to the Increment channel increments
	// the counter by 1, the value sent is ignored.
	//
	// Sending a channel to the Query channel causes
	// the value to be sent back on the channel.

	var counter Counter
	counter.Increment = make(chan int)
	counter.Query = make(chan chan int)

	go func() {
		value := 0

		for {
			select {

			case _ = <-counter.Increment:
				value++

			case reply_channel := <-counter.Query:
				reply_channel <- value
			}
		}
	}()

	return counter
}

func ServeCounterValue(counter Counter, host string, port string) {
	// Listens for tcp connwctions.  When a new
	// connection is made, writes the current counter
	// value as a string and closes the connection.

	lis, err := net.Listen("tcp", host+":"+port)
	if nil != err {
		fmt.Println("Error on listen:", err.Error())
		os.Exit(1)
	}
	defer lis.Close()
	fmt.Printf("Listening on %s:%s\n", LISTEN_HOST, LISTEN_PORT)

	for {
		conn, err := lis.Accept()
		if nil != err {
			fmt.Println("Error on accept: ", err.Error())
			os.Exit(1) // FIXME should not Exit() here
		}

		// serve the request (send the current value of the counter)
		go func() {
			c := make(chan int)
			counter.Query <- c
			value := <-c
			conn.Write([]byte(fmt.Sprintf("%d", value)))
			conn.Close()
		}()
	}
}

func main() {

	// GOAL : setup a tick-counter to track the number of pulses
	tick_counter := CreateTickCounter()

	// GOAL : accept network connections and serve up the counter value
	go ServeCounterValue(tick_counter, LISTEN_HOST, LISTEN_PORT)

	// GOAL: watch the pin forever
	watcher := gpio.NewWatcher()
	watcher.AddPin(INPUT_PIN)
	defer watcher.Close()
	fmt.Printf("watching GPIO%d\n", INPUT_PIN)

	for {
		pin, value := watcher.Watch()
		fmt.Printf("%s : (epoch %d) : read %d from gpio %d\n", time.Now(), time.Now().Unix(), value, pin)

		if 1 == value {
			tick_counter.Increment <- 0
		}
	}
}
