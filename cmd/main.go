package main

import (
	"Golang-Concurrency/pkg/publisher"
	"Golang-Concurrency/pkg/subscriber"
	"Golang-Concurrency/pkg/task"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	broadcaster := publisher.NewPublisher(ctx)

	spender1 := subscriber.NewSubscriber(ctx, "Spender1")
	spender2 := subscriber.NewSubscriber(ctx, "Spender2")

	go broadcaster.Update()

	spender1.Subscribe(broadcaster)
	spender2.Subscribe(broadcaster)
	go spender1.Update()
	go spender2.Update()

	go func() {
		tick := time.Tick(time.Second * 5)
		for {
			select {
			case ti := <-tick:
				log.Println(ti)
				t := task.NewTask("Register", uint64(1), uint64(100), ti.String())
				broadcaster.Publish(*t)
			case <-ctx.Done():
				return
			}
		}
	}()

	fmt.Scanln()
	cancel()
}
