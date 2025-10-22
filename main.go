package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"server_checker/workerpool"
	"syscall" // системные вызовы (SIGTERM, SIGINT)
	"time"
)

func main() {
	// загружавем конфигаруцию из файла yaml
	config, err := LoadConfig("config.yaml") // Передаем наш путь
	if err != nil {
		log.Fatalf("failed to load config: %v", err) // Если ошибка - фаталим
	}

	log.Printf("loaded configuration: %d workers, %d Urls, interval: %v", config.Workers_count, len(config.Urls), config.Interval) // логируем что мы загружаем
	// fmt.Println(config.Urls[:])

	results := make(chan workerpool.Result)
	workerPool := workerpool.New(config.Workers_count, config.Request_timeout, results)

	workerPool.Init() // запускаем воркеров

	go generateJobs(workerPool, config)
	go proccessResults(results)

	// обработка сигналов для Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit // блокируемся до получения сигнала

	log.Println("Shutting down gracefully...")
	workerPool.Stop() // останавливаем
	log.Println("Server checker stooped")
}

func proccessResults(results chan workerpool.Result) {
	defer close(results)
	for result := range results {
		fmt.Println(result.Info())
	}

}

// вывод в консоль

func generateJobs(wp *workerpool.Pool, config *Config) {
	ticker := time.NewTicker(config.Interval)
	defer ticker.Stop()

	for range ticker.C {
		for _, urlConfig := range config.Urls {
			wp.Push(workerpool.Job{
				URL:            urlConfig.Url,
				Name:           urlConfig.Name,
				ExpectedStatus: urlConfig.ExpectedStatus,
			})
		}
		log.Printf("Scheduled %d URLs for checking", len(config.Urls))

	}
}

// через указанное время добавляем url для проверки
