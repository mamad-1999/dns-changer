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
	"github.com/mamad-1999/dns-changer/constants"
	"github.com/mamad-1999/dns-changer/display"
	"github.com/mamad-1999/dns-changer/dns"
	"github.com/mamad-1999/dns-changer/utils"
)

func main() {
	homeDir, err := os.UserHomeDir()
	utils.HandleError(err, constants.ErrFindingHomeDir)

	configDir := filepath.Join(homeDir, constants.ConfigDir, constants.ConfigFile)
	configPath := filepath.Join(configDir, constants.ConfigFile)

	// Ensure the config directory exists
	err = config.EnsureConfigDir(configDir)
	utils.HandleError(err, constants.ErrCreatingConfigDir)

	// Download or validate the config.json
	err = config.ValidateConfigFile(configPath)
	utils.HandleError(err, constants.ErrHandlingConfigFile)

	dnsConfigs, err := config.LoadDnsConfigs(configPath)
	utils.HandleError(err, constants.ErrParsingDnsConfig)

	display.DisplayDnsOptions(dnsConfigs)

	reader := bufio.NewReader(os.Stdin)
	var choice int
	for {
		fmt.Print(constants.SelectDnsPrompt)
		input, err := reader.ReadString('\n')
		utils.HandleError(err, constants.ErrReadingInput)

		// Trim whitespace and newline characters
		input = strings.TrimSpace(input)

		// Check if input is a number
		choice, err = strconv.Atoi(input)
		utils.HandleError(err, constants.ErrInvalidInput)

		// Validate choice range
		if choice == 0 {
			color.Green(constants.SuccessExit)
			return
		}

		if choice < 1 || choice > len(dnsConfigs) {
			color.Red(constants.ErrInvalidChoice, len(dnsConfigs))
			continue
		}

		break
	}

	selectedConfig := dnsConfigs[choice-1]
	resolvContent := dns.BuildResolvContent(selectedConfig)

	err = dns.WriteToResolv(resolvContent)
	utils.HandleError(err, constants.ErrWritingToResolv)

	color.Green(constants.SuccessDnsChanged, selectedConfig.Name)
}
