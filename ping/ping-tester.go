package ping

import (
	"fmt"
	"os/exec"
	"strings"
)

func PingDns(server string) string {
	cmd := exec.Command("ping", "-c", "1", "-W", "2", server)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "unreachable"
	}

	outputStr := string(output)
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		if strings.Contains(line, "time=") {
			parts := strings.Split(line, "time=")
			return fmt.Sprintf("%s ", parts[1])
		}
	}
	return "unknown"
}
