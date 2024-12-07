package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mamad-1999/dns-changer/config"
	"github.com/mamad-1999/dns-changer/display"
	"github.com/mamad-1999/dns-changer/dns"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error finding home directory: %s", err)
		return
	}

	configDir := filepath.Join(homeDir, ".config", "dns-changer")
	configPath := filepath.Join(configDir, "config.json")

	// Ensure the config directory exists
	if err := config.EnsureConfigDir(configDir); err != nil {
		color.Red("Error creating config directory: %s", err)
		return
	}

	// Download or validate the config.json
	if err := config.ValidateConfigFile(configPath); err != nil {
		color.Red("Error handling config.json: %s", err)
		return
	}

	dnsConfigs, err := config.LoadDnsConfigs(configPath)
	if err != nil {
		color.Red("Error parsing DNS config: %s", err)
		return
	}

	display.DisplayDnsOptions(dnsConfigs)

	// Get user input
	var choice int
	fmt.Print("Select a DNS server by number: ")
	fmt.Scan(&choice)

	if choice == 0 {
		color.Green("Exiting the program.")
		return
	}

	if choice < 1 || choice > len(dnsConfigs) {
		color.Red("Invalid choice. Please run the program again.")
		return
	}

	// Build and apply the selected DNS configuration
	selectedConfig := dnsConfigs[choice-1]
	resolvContent := dns.BuildResolvContent(selectedConfig)

	if err := dns.WriteToResolv(resolvContent); err != nil {
		color.Red("Error writing to /etc/resolv.conf: %s", err)
		return
	}

	color.Green("Successfully changed DNS to %s", selectedConfig.Name)
}
