package service

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
)

type sensorService struct {
	l *zap.Logger
}

type SensorService interface {
	ProcessLog(filepath string)
}

// Used as check that service implements all of interface's methods.
var _ SensorService = (*sensorService)(nil)

func NewSensorService(
	l *zap.Logger,
) SensorService {
	return &sensorService{
		l: l,
	}
}

func (ss *sensorService) ProcessLog(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		ss.l.Fatal("Failed opening log", zap.Error(err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		ss.l.Fatal("Error scanning files", zap.Error(err))
	}
}
