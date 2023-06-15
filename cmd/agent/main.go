package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	mc "yandex-project/internal/metrics-collector"
)

const (
	timeout        = 1
	port           = "8080"
	pollInterval   = 2
	reportInterval = 10
)

func main() {
	collector := mc.NewCollector()
	pollInterval := reportInterval * time.Second
	for {
		time.Sleep(pollInterval)
		gaugeMetrics, counterMetrics := collector.GetMetrics()
		fmt.Println(collector)

		for ind, val := range gaugeMetrics {
			metricType := "gauge"
			metricName := ind
			metricValue := strconv.Itoa(int(val))
			PostMetrics(metricType, metricName, metricValue)
		}

		for ind, val := range counterMetrics {
			metricType := "counter"
			metricName := ind
			metricValue := strconv.Itoa(int(val))
			PostMetrics(metricType, metricName, metricValue)
		}
	}

	//requestDump, err := httputil.DumpRequest(request, true)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(string(requestDump))

}

func PostMetrics(metricType, metricName, metricValue string) {
	url := fmt.Sprintf("http://127.0.0.1:%s/update/%s/%s/%s", port, metricType, metricName, metricValue)
	client := http.Client{}
	client.Timeout = time.Second * timeout

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	request.Header.Set("Content-Type", "text/plain")
	_, err = client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
