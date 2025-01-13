package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend_total_hits",
			Help: "Number of get requests to /",
		},
		[]string{"value"},
	)
)

var (
	root = make(chan int, 1000)
)

func Ack(type string){
	
}

func Publish(){
	select {
		
	}
}

func BuildServer() error{

	prometheus.MustRegister(totalRequests)

	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe("localhost:22211", nil)
}