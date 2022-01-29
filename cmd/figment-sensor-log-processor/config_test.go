package main

import (
	"strings"
	"testing"
)

const configStringYaml = `
LogFilePath: ".log"
`

// TestUnmarshalFromYaml must pass when YAML config string is valid
func TestUnmarshalFromYaml(t *testing.T) {
	t.Run("Unmarshal YAML file", func(t *testing.T) {
		config, err := unmarshalFromYaml(strings.NewReader(configStringYaml))
		if err != nil {
			t.Errorf("unmarshalFromYaml(path) error = %v", err)
		}

		if config.LogFilePath != ".log" {
			t.Errorf("unmarshalFromYaml(path) invalid unmarshal")
		}
	})
}
