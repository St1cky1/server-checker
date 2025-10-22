package workerpool

import (
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

// структура worker содержит http клиент с настройками таймаута

func newWorker(timeout time.Duration) *worker {
	return &worker{
		&http.Client{
			Timeout: timeout,
		},
	}
}

// конструктор создаёт воркера с http клиентом, у которого установлен таймаут

func (w worker) process(j Job) Result {
	result := Result{
		URL:            j.URL,
		Name:           j.Name,
		ExpectedStatus: j.ExpectedStatus,
	}

	now := time.Now()

	resp, err := w.client.Get(j.URL)
	if err != nil {
		result.Error = err

		return result
	}

	result.StatusCode = resp.StatusCode
	result.ResponseTime = time.Since(now)

	defer resp.Body.Close()
	return result
}

// заполняем result (метрики, запросы, статус коды)
