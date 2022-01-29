package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

type Config struct {
	LogFilePath         string              `yaml:"LogFilePath"`
}

type SensorServiceConfig struct {
}

// NewConfigFromYaml Will try opening file defined in the path parameter
// and unmarshal it to Config struct
func NewConfigFromYaml(path string) (Config, error) {
	handle, err := os.Open(path)

	if err != nil {
		return Config{}, err
	}
	defer handle.Close()

	return unmarshalFromYaml(handle)
}

// unmarshalFromYaml Expects io.Reader which it will try unmarshalling
// as YAML file to Config struct
func unmarshalFromYaml(stream io.Reader) (Config, error) {
	bytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return Config{}, errors.New("error reading bytes from stream")
	}

	instance := &Config{}
	err = yaml.Unmarshal(bytes, instance)
	if err != nil {
		return Config{}, errors.New("couldn't unmarshal config to YAML file")
	}

	return *instance, nil
}
