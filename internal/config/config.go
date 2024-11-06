package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	OAuthToken       string   `json:"OAuthToken"`
	OrgNumber        string   `json:"OrgNumber"`
	Queues           []string `json:"Queues"`
	ConnectionString string   `json:"ConnectionString"`
}

func LoadFromJsonFile(configPath string, cfg *Config) error {
	f, err := os.OpenFile(configPath, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return fmt.Errorf("config file can not be opened: %s", err.Error())
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return fmt.Errorf("config file parsing error: %s", err.Error())
	}
	return nil
}
