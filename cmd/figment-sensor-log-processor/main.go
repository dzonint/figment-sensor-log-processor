package main

import (
	"figment-sensor-log-processor/pkg/service"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	const configPath = "./config/config.yaml"

	var (
		l             *zap.Logger
		err           error
		sensorService service.SensorService
		conf          Config
	)

	// Init logger.
	l, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal()
	}
	defer l.Sync()
	l.Info("Logger initialized...")

	// Init config.
	conf, err = NewConfigFromYaml(configPath)
	if err != nil {
		l.Error("Failed initializing config", zap.Error(err))
		os.Exit(1)
	}
	l.Info("Config initialized...")

	// Init service.
	sensorService = service.NewSensorService(l)
	if err = sensorService.ProcessLog(conf.LogFilePath); err != nil {
		l.Error("processing log failed", zap.Error(err))
	}
}
