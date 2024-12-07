package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mamad-1999/dns-changer/config"
	"github.com/mamad-1999/dns-changer/display"
	"github.com/mamad-1999/dns-changer/dns"
	"github.com/mamad-1999/dns-changer/utils"
)

func main() {
	homeDir, err := os.UserHomeDir()
	utils.HandleError(err, "Error finding home directory")

	configDir := filepath.Join(homeDir, ".config", "dns-changer")
	configPath := filepath.Join(configDir, "config.json")

	// Ensure the config directory exists
	if err := config.EnsureConfigDir(configDir); err != nil {
		utils.HandleError(err, "Error creating config directory")
		return
	}

	// Download or validate the config.json
	if err := config.ValidateConfigFile(configPath); err != nil {
		utils.HandleError(err, "Error handling config.json")
		return
	}

	dnsConfigs, err := config.LoadDnsConfigs(configPath)
	if err != nil {
		utils.HandleError(err, "Error parsing DNS config")
		return
	}

	display.DisplayDnsOptions(dnsConfigs)

	reader := bufio.NewReader(os.Stdin)
	var choice int
	for {
		fmt.Print("Select a DNS server by number (or 0 to exit): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			utils.HandleError(err, "Error reading input. Please try again.")
			continue
		}

		// Trim whitespace and newline characters
		input = strings.TrimSpace(input)

		// Check if input is a number
		choice, err = strconv.Atoi(input)
		if err != nil {
			utils.HandleError(err, "Invalid input. Please enter a valid number.")
			continue
		}

		// Validate choice range
		if choice == 0 {
			color.Green("Exiting the program.")
			return
		}

		if choice < 1 || choice > len(dnsConfigs) {
			color.Red("Invalid choice. Please select a number between 1 and %d, or 0 to exit.", len(dnsConfigs))
			continue
		}

		break
	}

	selectedConfig := dnsConfigs[choice-1]
	resolvContent := dns.BuildResolvContent(selectedConfig)

	if err := dns.WriteToResolv(resolvContent); err != nil {
		utils.HandleError(err, "Error writing to /etc/resolv.conf")
		return
	}

	color.Green("Successfully changed DNS to %s", selectedConfig.Name)
}
