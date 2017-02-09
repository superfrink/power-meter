package main

import (
	"fmt"
	"time"

	"github.com/brian-armstrong/gpio"
)

const (
	INPUT_PIN = 27
)

func main() {

	watcher := gpio.NewWatcher()
	watcher.AddPin(INPUT_PIN)
	defer watcher.Close()
	fmt.Printf("watching GPIO%d\n", INPUT_PIN)

	for {
		pin, value := watcher.Watch()
		fmt.Printf("%s : (epoch %d) : read %d from gpio %d\n", time.Now(), time.Now().Unix(), value, pin)
	}
}
