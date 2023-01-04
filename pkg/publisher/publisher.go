package publisher

import (
	"Golang-Concurrency/pkg/task"
	"context"
)

type Publisher struct {
	ctx              context.Context
	publishChannel   chan task.Task
	subscribeChannel chan chan<- task.Task
	subscribers      []chan<- task.Task
}

type IPublisher interface {
	Subscribe(sub chan<- task.Task)
	Publish(_task task.Task)
	Update()
}

func (p *Publisher) Subscribe(sub chan<- task.Task) {
	// 구독 채널로 전송
	p.subscribeChannel <- sub
}

func (p *Publisher) Publish(_task task.Task) {
	// 배포 채널로 전송
	p.publishChannel <- _task
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeChannel:
			p.subscribers = append(p.subscribers, sub)
		case _task := <-p.publishChannel:
			for _, sub := range p.subscribers {
				sub <- _task
			}
		case <-p.ctx.Done():
			return
		}
	}
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{ctx: ctx, publishChannel: make(chan task.Task), subscribeChannel: make(chan chan<- task.Task), subscribers: make([]chan<- task.Task, 0)}
}
