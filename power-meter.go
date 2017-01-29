package main

import (
	"fmt"
	"github.com/brian-armstrong/gpio"
)

func main() {

	watcher := gpio.NewWatcher()
	watcher.AddPin(27)
	defer watcher.Close()

	go func() {
		for {
			pin, value := watcher.Watch()
			fmt.Printf("read %d from gpio %d\n", value, pin)
		}
	}()

}
