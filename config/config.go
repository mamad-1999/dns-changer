package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const configURL = "https://raw.githubusercontent.com/mamad-1999/dns-changer/refs/heads/master/config.json"

type DnsConfig struct {
	Name    string   `json:"name"`
	Servers []string `json:"servers"`
}

func EnsureConfigDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

func ValidateConfigFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("config.json not found locally. Downloading from GitHub...")
		return DownloadConfig(path)
	}
	return nil
}

func DownloadConfig(path string) error {
	resp, err := http.Get(configURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download config: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, body, 0644)
}

func LoadDnsConfigs(path string) ([]DnsConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dnsConfigs []DnsConfig
	if err := json.Unmarshal(data, &dnsConfigs); err != nil {
		return nil, err
	}

	return dnsConfigs, nil
}
