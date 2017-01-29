package main

import (
	"fmt"
	"time"

	"github.com/brian-armstrong/gpio"
)

func main() {

	watcher := gpio.NewWatcher()
	watcher.AddPin(27)
	defer watcher.Close()

	go func() {
		for {
			pin, value := watcher.Watch()
			fmt.Printf("%d : read %d from gpio %d\n", time.Now().Unix(), value, pin)
		}
	}()
}
