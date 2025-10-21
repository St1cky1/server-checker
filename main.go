package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"server_checker/workerpool"
)

const (
	INTERVAL        = time.Second * 10
	REQUEST_TIMEOUT = time.Second * 2
	WORKERS_COUNT   = 3
)

var urls = []string{
	"https://github.com/St1cky1",
	"https://books.yandex.ru/?from=tableau_yabro",
	"https://sokolov.ru/?utm_referrer=https%3A%2F%2Fwww.yandex.ru%2Fclck%2Fjsredir%3Ffrom%3Dyandex.ru%3Bsuggest%3Bbrowser%26text%3D",
	"https://google.com/",
	"https://golang.org/",
}

func main() {
	results := make(chan workerpool.Result)
	workerPool := workerpool.New(WORKERS_COUNT, REQUEST_TIMEOUT, results)

	workerPool.Init()

	go generateJobs(workerPool)
	go proccessResults(results)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	workerPool.Stop()
}

func proccessResults(results chan workerpool.Result) {
	go func() {
		for result := range results {
			fmt.Println(result.Info())
		}
	}()
}

func generateJobs(wp *workerpool.Pool) {
	for {
		for _, url := range urls {
			wp.Push(workerpool.Job{URL: url})
		}

		time.Sleep(INTERVAL)
	}
}
