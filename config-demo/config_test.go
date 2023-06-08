package config_demo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	PrintConfig()
}

func TestConfig2(t *testing.T) {
	PrintTomlFromFile()
}

type MyConfig struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJson(t *testing.T) {
	jsonString := `{"name": "Alice", "age": 18}`

	var config MyConfig
	err := json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", config.Name, config.Age)
}
