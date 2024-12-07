package dns

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mamad-1999/dns-changer/config"
)

func BuildResolvContent(config config.DnsConfig) string {
	var resolvContent strings.Builder
	resolvContent.WriteString("# Generated by Dns-changer\n")
	for _, server := range config.Servers {
		resolvContent.WriteString(fmt.Sprintf("nameserver %s\n", server))
	}
	return resolvContent.String()
}

func WriteToResolv(content string) error {
	cmd := exec.Command("sudo", "sh", "-c", fmt.Sprintf("echo '%s' > /etc/resolv.conf", content))
	return cmd.Run()
}
