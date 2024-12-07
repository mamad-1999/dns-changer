package ping

import (
	"os/exec"
	"strings"
)

type PingResult struct {
	Time   string // e.g., "23.4 ms"
	Status string // e.g., "reachable", "unreachable", "unknown"
}

func PingDns(servers []string) []PingResult {
	var results []PingResult

	for _, server := range servers {
		cmd := exec.Command("ping", "-c", "1", "-W", "1", server)
		output, err := cmd.CombinedOutput()

		var pingRes PingResult
		if err != nil {
			pingRes = PingResult{
				Time:   "N/A",
				Status: "unreachable",
			}
		} else {
			outputStr := string(output)
			lines := strings.Split(outputStr, "\n")
			for _, line := range lines {
				if strings.Contains(line, "time=") {
					parts := strings.Split(line, "time=")
					pingRes = PingResult{
						Time:   parts[1],
						Status: "reachable",
					}
				}
			}
			if pingRes.Status == "" {
				pingRes = PingResult{
					Time:   "N/A",
					Status: "unknown",
				}
			}
		}

		results = append(results, pingRes)
	}

	return results
}
