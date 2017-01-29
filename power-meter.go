package main

import (
	"fmt"
	"time"

	"github.com/brian-armstrong/gpio"
)

func main() {

	fmt.Printf("time: %d\n", time.Now().Unix())
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
