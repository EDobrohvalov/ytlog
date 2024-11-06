package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	OAuthToken       string   `json:"OAuthToken"`
	OrgNumber        string   `json:"OrgNumber"`
	Queues           []string `json:"Queues"`
	ConnectionString string   `json:"ConnectionString"`
}

func LoadFromJsonFile(configPath string, cfg *Config) error {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("config file can not be opened: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue[3:], &cfg)
	if err != nil {
		return fmt.Errorf("config file parsing error: %s", err.Error())
	}
	return nil
}
