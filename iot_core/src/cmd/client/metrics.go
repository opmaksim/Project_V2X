package main

import (
	"os"
	"os/signal"
	service "project/v2x/iot-core/internal/services"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	monitor := &service.SystemMetricsProcessor{}
	metrics := monitor.NewSystemMetricsProcessor()
	go metrics.ExecuteMonitoring()

	// for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// wait shutdown when occure signal via sigChan
	select {
	case sig := <-sigChan:
		log.Infof("Receive signal: %s, Shutting down...\n", sig)
	}
}
