package workerpool

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Job struct {
	Name           string // читаемое имя
	URL            string
	ExpectedStatus int // ожидаемый статус
}

type Result struct {
	Name           string
	URL            string
	ExpectedStatus int
	StatusCode     int
	ResponseTime   time.Duration
	Error          error
}

// результаты выполнения с метриками

func (r Result) Info() string {
	if r.Error != nil {
		return fmt.Sprintf("[ERROR] - %s [%s] - %s", r.Name, r.URL, r.Error.Error())
	}

	// Проверяем, совпадает ли ожидаемый статус с полученным
	status := "SUCCESS"
	if r.StatusCode != r.ExpectedStatus {
		status = "WRONG STATUS"
	}

	return fmt.Sprintf("[%s] - [%s] - Status: %d (expected status: %d), Time: %s", status, r.Name, r.StatusCode, r.ExpectedStatus, r.ResponseTime.String())
}

// форматирование ответов

type Pool struct {
	worker       *worker
	workersCount int

	jobs    chan Job    // канал для задач
	results chan Result // канал для результатов

	wg      *sync.WaitGroup // синхрон
	stopped bool            // флаг для остановки
}

// управление работой воркеров

func New(workersCount int, timeout time.Duration, results chan Result) *Pool {
	return &Pool{
		worker:       newWorker(timeout),
		workersCount: workersCount,
		jobs:         make(chan Job),
		results:      results,
		wg:           new(sync.WaitGroup),
	}
}

func (p *Pool) Init() {
	for i := 0; i < p.workersCount; i++ {
		go p.initWorker(i)
	}
}

// запускаем указанное количество воркеров

func (p *Pool) Push(j Job) {
	if p.stopped {
		return
	}

	p.jobs <- j
	p.wg.Add(1)
}

// добавляем задачу в пул

func (p *Pool) Stop() {
	p.stopped = true
	close(p.jobs)
	p.wg.Wait()
}

// дожидаемся завершения текущих задач и завершаем выполнение

func (p *Pool) initWorker(id int) {
	for job := range p.jobs {
		time.Sleep(time.Second)
		p.results <- p.worker.process(job)
		p.wg.Done()
	}

	log.Printf("[worker ID %d] finished proccesing", id)
}

// цикл работы воркеров
