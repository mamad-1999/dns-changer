package ping

import (
	"os/exec"
	"strings"
)

type PingResult struct {
	Time   string // e.g., "23.4 ms"
	Status string // e.g., "reachable", "unreachable", "unknown"
}

func PingDns(server string) PingResult {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", server)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return PingResult{
			Time:   "N/A",
			Status: "unreachable",
		}
	}

	outputStr := string(output)
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		if strings.Contains(line, "time=") {
			parts := strings.Split(line, "time=")
			return PingResult{
				Time:   parts[1],
				Status: "reachable",
			}
		}
	}
	return PingResult{
		Time:   "N/A",
		Status: "unknown",
	}
}
