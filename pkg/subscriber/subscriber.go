package subscriber

import (
	"Golang-Concurrency/pkg/publisher"
	"Golang-Concurrency/pkg/task"
	"context"
	"log"
)

type Subscriber struct {
	ctx     context.Context
	name    string
	channel chan task.Task
}

func (s *Subscriber) Subscribe(pub *publisher.Publisher) {
	pub.Subscribe(s.channel)
}

func (s *Subscriber) Update() {
	for {
		select {
		case _task := <-s.channel:
			log.Println("Task is updated in ", s.name, _task)
		case <-s.ctx.Done():
			return
		}
	}
}

func NewSubscriber(ctx context.Context, name string) ISubscriber {
	return &Subscriber{ctx: ctx, name: name, channel: make(chan task.Task)}
}

type ISubscriber interface {
	Subscribe(pub *publisher.Publisher)
	Update()
}
