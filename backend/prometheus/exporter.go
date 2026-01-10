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
	totalLoad = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend_total_load",
			Help: "Number of get requests to /load that get puzzle from MongoDb",
		},
		[]string{"value"},
	)
	totalSeen = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend_total_seen",
			Help: "Number of seen on puzzles",
		},
		[]string{"value"},
	)
	totalSolved = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend_total_solved",
			Help: "Number of people solved the puzles",
		},
		[]string{"value"},
	)
	numberOfUniqueUsers = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "number_of_unique_users",
			Help: "Number of unique users that visited the site",
		},
		[]string{"value"},
	)
	grandSolveTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "number_of_all_solves",
			Help: "Number of all solves of all puzzles together",
		},
		[]string{"value"},
	)
)

var (
	root = make(chan int, 1000)
	load = make(chan int, 1000)
	seen = make(chan int, 1000)
	solved = make(chan int, 1000)
	unique = make(chan int)
	grandsolve = make(chan int)
)

func Publish(ack_type string){
	switch ack_type {
	case "root":
		root <- 1
	case "load":
		load <- 1
	case "seen":
		seen <- 1
	case "solved":
		solved <- 1
	}
}

func PublishWithValue(ack_type string, aa int){
	switch ack_type {
	case "unique":
		unique <- aa
	case "grandsolve":
		grandsolve <- aa
	}
}

func ExecLoop(){
	for{
		select {
		case  _= <- root:
			totalRequests.WithLabelValues("value").Inc()
		case  _= <- load:
			totalLoad.WithLabelValues("value").Inc()
		case  _= <- seen:
			totalSeen.WithLabelValues("value").Inc()
		case  _= <- solved:
			totalSolved.WithLabelValues("value").Inc()
		case  un := <- unique:
			numberOfUniqueUsers.WithLabelValues("value").Set(float64(un))
		case  gs := <- grandsolve:
			grandSolveTotal.WithLabelValues("value").Set(float64(gs))
		}
	}
}

func BuildServer() error{

	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(totalLoad)
	prometheus.MustRegister(totalSeen)
	prometheus.MustRegister(totalSolved)
	prometheus.MustRegister(numberOfUniqueUsers)
	prometheus.MustRegister(grandSolveTotal)

	go ExecLoop()

	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe("localhost:22211", nil)
}