package task

type Task struct {
	name       string
	tokenId    uint64
	price      uint64
	saleMethod string
}

func NewTask(name string, tokenId uint64, price uint64, saleMethod string) *Task {
	return &Task{name: name, tokenId: tokenId, price: price, saleMethod: saleMethod}
}
