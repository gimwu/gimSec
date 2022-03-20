package main

import (
	"fmt"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(arg_f func() error) *Task {
	t := Task{
		f: arg_f,
	}
	return &t
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChannel chan *Task
	JobsChannel  chan *Task
	worker_num   int
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		worker_num:   cap,
	}
	return &p
}

func (p *Pool) worker(worker_ID int) {
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("worker ID: ", worker_ID, "finished")
	}
}

func (p *Pool) run() {
	for i := 0; i < p.worker_num; i++ {
		go p.worker(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}

func main() {
	task := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	p := NewPool(4)

	go func() {
		for i := 0; i < 10; i++ {
			p.EntryChannel <- task
		}
	}()
	p.run()
}
