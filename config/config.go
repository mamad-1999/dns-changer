package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mamad-1999/dns-changer/utils"
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
	utils.HandleError(err, "Failed to download config")

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to download config: %s", resp.Status)
		utils.HandleError(err, "Failed to download config")
	}

	body, err := ioutil.ReadAll(resp.Body)
	utils.HandleError(err, "Failed to read response body")

	return ioutil.WriteFile(path, body, 0644)
}

func LoadDnsConfigs(path string) ([]DnsConfig, error) {
	data, err := ioutil.ReadFile(path)
	utils.HandleError(err, "Failed to read DNS config file")

	var dnsConfigs []DnsConfig
	err = json.Unmarshal(data, &dnsConfigs)
	utils.HandleError(err, "Failed to parse DNS config file")

	return dnsConfigs, nil
}
