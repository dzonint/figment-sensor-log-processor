package service

import (
	"bufio"
	"errors"
	"figment-sensor-log-processor/pkg/sensors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
)

const (
	referenceKeyword     = "reference"
	stringSeparator      = " "
	sensorLineWordCount  = 2
	readingLineWordCount = 3
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
		ss.l.Fatal("failed opening log", zap.Error(err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Scan first line outside of loop to collect reference values.
	scanner.Scan()
	temp, humidity, err := ss.determineTemperatureAndHumidity(scanner.Text())
	if err != nil {
		ss.l.Fatal("failed determining temperature and humidity", zap.Error(err))
	}

	var sensor *sensors.Sensor
	var sensorType, name string
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), stringSeparator)
		lineWordCount := len(data)
		if lineWordCount == sensorLineWordCount {
			if sensorType != "" { // This means that non-first sensor is encountered so we print the current sensor's state.
				sensor.PrintState()
			}
			sensorType, name = data[0], data[1]
			sensor = sensors.NewSensor(sensorType, name, temp, humidity)
			continue
		} else if lineWordCount == readingLineWordCount {
			if value, err := strconv.ParseFloat(data[2], 64); err != nil {
				ss.l.Error("failed to parse value", zap.Error(err))
			} else {
				sensor.Values = append(sensor.Values, value)
			}
		}
	}
	// One last sensor print after it reaches EOF.
	sensor.PrintState()

	if err := scanner.Err(); err != nil {
		ss.l.Fatal("error scanning files", zap.Error(err))
	}
}

func (ss *sensorService) determineTemperatureAndHumidity(line string) (temp float64, humidity float64, err error) {
	data := strings.Split(line, stringSeparator)
	if len(data) != 3 {
		return 0, 0, errors.New("insufficient number of arguments")
	} else if data[0] != referenceKeyword {
		return 0, 0, errors.New("could not find keyword 'reference'")
	} else if temp, err = strconv.ParseFloat(data[1], 64); err != nil {
		return 0, 0, errors.New(fmt.Sprintf("failed to parse temperature parameter: %v", err.Error()))
	} else if humidity, err = strconv.ParseFloat(data[2], 64); err != nil {
		return 0, 0, errors.New(fmt.Sprintf("failed to parse humidity parameter: %v", err.Error()))
	}
	return temp, humidity, nil
}
