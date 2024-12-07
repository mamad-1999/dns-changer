package display

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/mamad-1999/dns-changer/config"
	"github.com/mamad-1999/dns-changer/ping"
	"github.com/rodaine/table"
)

func DisplayDnsOptions(dnsConfigs []config.DnsConfig) {
	t := table.New("", "DNS Server", "Ping Time")
	color.Yellow("I am pinging the DNS servers concurrently, please wait...")

	// Channel to collect ping results
	type PingResult struct {
		Index      int
		ServerName string
		PingTime   string
		Status     string
	}

	results := make(chan PingResult, len(dnsConfigs))
	var wg sync.WaitGroup

	// Launch Goroutines for concurrent pinging
	for i, config := range dnsConfigs {
		wg.Add(1)
		go func(i int, serverName string, serverAddress string) {
			defer wg.Done()
			pingRes := ping.PingDns(serverAddress)
			results <- PingResult{
				Index:      i + 1,
				ServerName: serverName,
				PingTime:   pingRes.Time,
				Status:     pingRes.Status,
			}
		}(i, config.Name, config.Servers[0]) // Assume the first server is used for pinging
	}

	// Close the channel when all Goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and display results
	for result := range results {
		var pingDisplay string

		switch result.Status {
		case "reachable":
			pingDisplay = color.GreenString(result.PingTime)
		case "unreachable":
			pingDisplay = color.RedString("unreachable")
		case "unknown":
			pingDisplay = color.YellowString("unknown")
		}

		t.AddRow(
			fmt.Sprintf("%d", result.Index),
			fmt.Sprintf("%-15s", result.ServerName),
			fmt.Sprintf("%-20s", pingDisplay),
		)
	}

	t.Print()
	fmt.Println("0. Exit")
}
