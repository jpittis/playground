package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "app_counter_total",
	Help: "A default counter that keeps going up",
})

func main() {
	go func() {
		log.Info("Starting counter loop...")
		for {
			counter.Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	log.Info("Listening...")
	log.Fatal(http.ListenAndServe(":2112", nil))
}
