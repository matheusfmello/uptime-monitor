package jobs

import (
	"log"
	"net/http"
	"time"

	"github.com/matheusfmello/uptime-monitor/internal/services"
)

type MonitorChecker struct {
	service services.MonitorServiceInterface
}

func NewMonitorChecker(ms services.MonitorServiceInterface) *MonitorChecker {
	return &MonitorChecker{ms}
}

func (mc *MonitorChecker) CheckMonitors() {
	monitors, err := mc.service.GetDueMonitors()
	if err != nil {
		log.Printf("Error fetching monitors: %v", err)
		return
	}

	for _, monitor := range monitors {
		go func(monitor services.MonitorInfo) {
			status := "UP"
			start := time.Now()
			resp, err := http.Get(monitor.URL)
			duration := time.Since(start)

			if err != nil || resp.StatusCode > 400 {
				status = "DOWN"
			}

			log.Printf("Checked %s: %s (%d)", monitor.URL, status, duration.Milliseconds())

			mc.service.SaveCheckResult(monitor.ID, status, duration.Milliseconds())
		}(monitor)

	}

}

func (mc *MonitorChecker) Run() {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		log.Println("Running monitor checker...")
		mc.CheckMonitors()
	}
}
