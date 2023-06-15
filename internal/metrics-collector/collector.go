package metricscollector

import (
	"math/rand"
	"runtime"
)

type gauge float64
type counter int64

type Metrics struct {
	gaugeMetrics   map[string]gauge
	counterMetrics map[string]counter
}

func NewCollector() Metrics {
	return Metrics{
		gaugeMetrics:   make(map[string]gauge),
		counterMetrics: make(map[string]counter),
	}
}

func (m *Metrics) GetMetrics() (map[string]gauge, map[string]counter) {
	rand.Int()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	m.counterMetrics["PollCount"]++
	m.gaugeMetrics["Alloc"] = gauge(rtm.Alloc)
	m.gaugeMetrics["BuckHashSys"] = gauge(rtm.BuckHashSys)
	m.gaugeMetrics["Frees"] = gauge(rtm.Frees)
	m.gaugeMetrics["GCCPUFraction"] = gauge(rtm.GCCPUFraction)
	m.gaugeMetrics["GCSys"] = gauge(rtm.GCSys)
	m.gaugeMetrics["HeapAlloc"] = gauge(rtm.HeapAlloc)
	m.gaugeMetrics["HeapIdle"] = gauge(rtm.HeapIdle)
	m.gaugeMetrics["HeapInuse"] = gauge(rtm.HeapInuse)
	m.gaugeMetrics["HeapObjects"] = gauge(rtm.HeapObjects)
	m.gaugeMetrics["HeapReleased"] = gauge(rtm.HeapReleased)
	m.gaugeMetrics["HeapSys"] = gauge(rtm.HeapSys)
	m.gaugeMetrics["LastGC"] = gauge(rtm.LastGC)
	m.gaugeMetrics["Lookups"] = gauge(rtm.Lookups)
	m.gaugeMetrics["MCacheInuse"] = gauge(rtm.MCacheInuse)
	m.gaugeMetrics["MCacheSys"] = gauge(rtm.MCacheSys)
	m.gaugeMetrics["MSpanInuse"] = gauge(rtm.MSpanInuse)
	m.gaugeMetrics["MSpanSys"] = gauge(rtm.MSpanSys)
	m.gaugeMetrics["Mallocs"] = gauge(rtm.Mallocs)
	m.gaugeMetrics["NextGC"] = gauge(rtm.NextGC)
	m.gaugeMetrics["NumForcedGC"] = gauge(rtm.NumForcedGC)
	m.gaugeMetrics["NumGC"] = gauge(rtm.NumGC)
	m.gaugeMetrics["OtherSys"] = gauge(rtm.OtherSys)
	m.gaugeMetrics["PauseTotalNs"] = gauge(rtm.PauseTotalNs)
	m.gaugeMetrics["StackInuse"] = gauge(rtm.StackInuse)
	m.gaugeMetrics["StackSys"] = gauge(rtm.StackSys)
	m.gaugeMetrics["Sys"] = gauge(rtm.Sys)
	m.gaugeMetrics["TotalAlloc"] = gauge(rtm.TotalAlloc)
	m.gaugeMetrics["RandomValue"] = gauge(rand.Int())

	return m.gaugeMetrics, m.counterMetrics
}
