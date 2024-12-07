package display

import (
	"fmt"
	"sort"
	"sync"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mamad-1999/dns-changer/config"
	"github.com/mamad-1999/dns-changer/ping"
)

func DisplayDnsOptions(dnsConfigs []config.DnsConfig) {
	// Inform the user that DNS servers are being pinged
	color.Yellow("I am pinging the DNS servers concurrently, please wait...")

	// Channel to collect ping results
	type PingResult struct {
		Index      int
		ServerName string
		Server1    string
		PingTime1  string
		Server2    string
		PingTime2  string
		Status1    string
		Status2    string
	}

	results := make(chan PingResult, len(dnsConfigs))
	var wg sync.WaitGroup

	// Launch Goroutines for concurrent pinging
	for i, config := range dnsConfigs {
		wg.Add(1)
		go func(i int, serverName string, servers []string) {
			defer wg.Done()
			// Ping both servers in the DNS config
			pingRes := ping.PingDns(servers)

			// Add both ping results to the channel
			results <- PingResult{
				Index:      i + 1,
				ServerName: serverName,
				Server1:    servers[0],
				PingTime1:  pingRes[0].Time,
				Status1:    pingRes[0].Status,
				Server2:    servers[1],
				PingTime2:  pingRes[1].Time,
				Status2:    pingRes[1].Status,
			}
		}(i, config.Name, config.Servers)
	}

	// Close the channel when all Goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Create the table
	t := table.NewWriter()
	t.SetOutputMirror(color.Output)
	t.AppendHeader(table.Row{"IN", "DNS Server", "Server 1 (Ping)", "Server 2 (Ping)"})

	// Collect and display results
	var resultSlice []PingResult
	for result := range results {
		resultSlice = append(resultSlice, result)
	}

	// Sort the results by the original index
	sort.Slice(resultSlice, func(i, j int) bool {
		return resultSlice[i].Index < resultSlice[j].Index
	})

	// Add the rows to the table in the correct order
	for _, result := range resultSlice {
		// Format the ping results for both servers
		var pingDisplay1, pingDisplay2 string

		// Handle server 1 ping display
		switch result.Status1 {
		case "reachable":
			pingDisplay1 = color.GreenString(result.PingTime1)
		case "unreachable":
			pingDisplay1 = color.RedString("unreachable")
		case "unknown":
			pingDisplay1 = color.YellowString("unknown")
		}

		// Handle server 2 ping display
		switch result.Status2 {
		case "reachable":
			pingDisplay2 = color.GreenString(result.PingTime2)
		case "unreachable":
			pingDisplay2 = color.RedString("unreachable")
		case "unknown":
			pingDisplay2 = color.YellowString("unknown")
		}

		// Add the row to the table
		t.AppendRow(table.Row{
			fmt.Sprintf("%d", result.Index),
			result.ServerName,
			pingDisplay1,
			pingDisplay2,
		})
	}

	t.SetStyle(table.StyleLight)
	t.Render()
	fmt.Println("0. Exit")
}
