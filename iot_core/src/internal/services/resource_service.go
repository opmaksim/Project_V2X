package service

import (
	"project/v2x/iot-core/pkg/monitor"
	"time"

	log "github.com/sirupsen/logrus"
)

type SystemMetricsProcessor struct {
	metrics *monitor.SystemMetrics
}

func (__this *SystemMetricsProcessor) NewSystemMetricsProcessor() *SystemMetricsProcessor {
	return &SystemMetricsProcessor{
		metrics: &monitor.SystemMetrics{}, // ✅ 올바르게 초기화
	}
}

func (__this *SystemMetricsProcessor) ExecuteMonitoring() {
	log.Info("Start monitoring system metrics...")

	for {
		metrics, err := __this.metrics.GetSystemMetrics()
		if err != nil {
			log.Errorf("Error during monitoring: %v", err)
			continue
		}

		log.Info("==============================")
		log.Infof("CPU 사용량: %.2f%%\n", metrics.CPUUsage)
		log.Infof("메모리 사용량: %.2f%%\n", metrics.MemoryUsage)
		log.Infof("디스크 사용량: %.2f%%\n", metrics.DiskUsage)
		log.Infof("수신 속도: %d bps / 송신 속도 %d bps\n", metrics.NetRecvSpeed, metrics.NetSentSpeed)

		time.Sleep(30 * time.Second)
	}
}
