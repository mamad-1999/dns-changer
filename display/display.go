package display

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mamad-1999/dns-changer/config"
	"github.com/mamad-1999/dns-changer/ping"
	"github.com/rodaine/table"
)

func DisplayDnsOptions(dnsConfigs []config.DnsConfig) {
	t := table.New("", "DNS Server", "Ping Time")
	color.Yellow("I am pinging the DNS server, please wait 2 seconds...")

	for i, config := range dnsConfigs {
		pingResult := ping.PingDns(config.Servers[0])
		t.AddRow(fmt.Sprintf("%d", i+1), fmt.Sprintf("%-15s", config.Name), fmt.Sprintf("%-20s", pingResult))
	}

	t.Print()
	fmt.Println("0. Exit")
}
