package service

import (
	"bufio"
	"errors"
	"figment-sensor-log-processor/pkg/sensors"
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
	fileEmptyMessage     = "opened file is empty"
)

var failedOpeningLogError = errors.New("failed opening log")
var determineTemperatureAndHumidityError = errors.New("failed determining temperature and humidity")
var scannerError = errors.New("error scanning files")
var notEnoughArgumentsError = errors.New("insufficient number of arguments")
var missingReferenceKeywordError = errors.New("could not find keyword 'reference'")
var failedToParseTemperatureError = errors.New("failed to parse temperature parameter")
var failedToParseHumidityError = errors.New("failed to parse humidity parameter")

type sensorService struct {
	l *zap.Logger
}

type SensorService interface {
	ProcessLog(filepath string) (err error)
	determineTemperatureAndHumidity(line string) (temp float64, humidity float64, err error)
	isFileEmpty(file *os.File) bool
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

func (ss *sensorService) ProcessLog(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		ss.l.Error(failedOpeningLogError.Error(), zap.Error(err))
		return failedOpeningLogError
	}
	defer file.Close()

	if ss.isFileEmpty(file) {
		ss.l.Info(fileEmptyMessage)
		return nil
	}

	scanner := bufio.NewScanner(file)
	// Scan first line outside of loop to collect reference values.
	scanner.Scan()
	temp, humidity, err := ss.determineTemperatureAndHumidity(scanner.Text())
	if err != nil {
		ss.l.Error(determineTemperatureAndHumidityError.Error(), zap.Error(err))
		return determineTemperatureAndHumidityError
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
		} else if lineWordCount == readingLineWordCount {
			if value, err := strconv.ParseFloat(data[2], 64); err != nil {
				ss.l.Error("failed to parse value", zap.Error(err))
				ss.l.Warn("invalid value encountered! Final evaluation might not be accurate", zap.String("data", scanner.Text()))
			} else {
				sensor.Values = append(sensor.Values, value)
			}
		}
	}
	// One last sensor print after it reaches EOF.
	sensor.PrintState()

	if err = scanner.Err(); err != nil {
		ss.l.Error(scannerError.Error(), zap.Error(err))
		return scannerError
	}
	return nil
}

// determineTemperatureAndHumidity takes the first line of .log and determines reference temperature and humidity.
// Example: 'reference 75.0 45.0'
func (ss *sensorService) determineTemperatureAndHumidity(line string) (temp float64, humidity float64, err error) {
	data := strings.Split(line, stringSeparator)
	if len(data) != 3 {
		return 0, 0, notEnoughArgumentsError
	} else if data[0] != referenceKeyword {
		return 0, 0, missingReferenceKeywordError
	} else if temp, err = strconv.ParseFloat(data[1], 64); err != nil {
		return 0, 0, failedToParseTemperatureError
	} else if humidity, err = strconv.ParseFloat(data[2], 64); err != nil {
		return 0, 0, failedToParseHumidityError
	}
	return temp, humidity, nil
}

func (ss *sensorService) isFileEmpty(file *os.File) bool {
	fi, err := file.Stat()
	if err != nil {
		return true
	}
	return fi.Size() == 0
}
