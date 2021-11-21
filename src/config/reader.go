package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Configuration struct {
	MaxConnections int `json:"maxConnections"`
	MaxArgs int `json:"maxArgs"`
	MaxRequests int `json:"maxRequests"`
}

func Read() Configuration {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file, err := os.Open(path + "/../cfg/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var config Configuration
	json.Unmarshal(byteValue, &config)

	return config
}
