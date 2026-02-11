package http

import (
	"log"
	"time"
)

type Task struct {
	Type    string
	Payload interface{}
}

type TaskQueue struct {
	ch chan Task
}

func NewTaskQueue(size int) *TaskQueue {
	q := &TaskQueue{ch: make(chan Task, size)}
	go q.worker()
	return q
}

func (q *TaskQueue) Enqueue(t Task) {
	select {
	case q.ch <- t:
	default:
	}
}

func (q *TaskQueue) worker() {
	for t := range q.ch {
		log.Printf("processing task: %s", t.Type)
		time.Sleep(10 * time.Millisecond)
	}
}
