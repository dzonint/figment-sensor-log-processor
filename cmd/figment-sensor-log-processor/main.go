package main

import (
	"figment-sensor-log-processor/pkg/service"
	"go.uber.org/zap"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func getProjectRootDirectoryPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	position := strings.LastIndex(basepath, "pkg")
	return basepath[:position]
}

func main() {
	var (
		l             *zap.Logger
		err           error
		sensorService service.SensorService
	)

	// Logger.
	l, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal()
	}
	defer l.Sync()
	l.Info("Logger initialized...")

	// Service.
	sensorService = service.NewSensorService(l)
	sensorService.ProcessLog(getProjectRootDirectoryPath() + ".log")
}
