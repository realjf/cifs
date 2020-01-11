package config

import "testing"

func TestNewConfig(t *testing.T) {
	config := NewConfig().LoadConfig("./config.json")

	t.Fatal(config.Data)
}

