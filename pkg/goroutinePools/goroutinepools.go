package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

type Job struct {
	W      http.ResponseWriter
	R      *http.Request
	Result chan map[string]interface{}
}

var (
	MaxQueue   = 1000
	MaxWorkers int
	JobQueue   chan Job
)

type Worker struct {
	WorkerPool chan chan Job
	JobChan    chan Job
	Quit       chan bool
}

func NewWorker(workPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workPool,
		JobChan:    make(chan Job),
		Quit:       make(chan bool),
	}
}

func JobHandler(j Job) {

	fmt.Println(j.R.Method)
	j.Result <- map[string]interface{}{"name": "王晓勃"}
	defer close(j.Result)
	// j.W.Header().Set("Content-Type", "application/json")
	// j.W.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]string{"name": "王晓勃"})
	return
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChan
			select {
			case job := <-w.JobChan:
				go JobHandler(job)
			case <-w.Quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

type Dispatcher struct {
	MaxWorkers int
	WorkerPool chan chan Job
	Quit       chan bool
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		MaxWorkers: maxWorkers,
		WorkerPool: make(chan chan Job, maxWorkers),
		Quit:       make(chan bool),
	}
}

func (d *Dispatcher) Run() {
	// 启动指定worker
	for i := 0; i < MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func (d *Dispatcher) Stop() {
	go func() {
		d.Quit <- true
	}()
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(j Job) {
				jobChan := <-d.WorkerPool
				jobChan <- j
			}(job)
		case <-d.Quit:
			return
		}
	}
}

func init() {
	MaxWorkers = runtime.NumCPU()
	JobQueue = make(chan Job, MaxQueue)

	d := NewDispatcher(MaxWorkers)
	d.Run()
}

func handler(w http.ResponseWriter, r *http.Request) {
	job := Job{
		W:      w,
		R:      r,
		Result: make(chan map[string]interface{}),
	}
	JobQueue <- job

	res := <-job.Result
	w.Header().Set("Content-Type", "application/json")
	// j.W.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func main() {
	addr := "0.0.0.0:8090"
	log.Println("service start, listening on", addr)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
