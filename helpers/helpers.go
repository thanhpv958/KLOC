package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type AuthorStats struct {
	Author  string
	Commits int
	Ins     int
	Del     int
}

func GetGitData() []AuthorStats {
	cmd := exec.Command("bash", "-c", `cd /home/thanhpv958/Documents/dev/fabbi/sp/rage-esports-repo && git log --no-merges --pretty=format:%an --numstat --since=2022-01-01 --until=2023-12-31 | awk '/./ && !author { author = $0; next } author { commits[author]++; ins[author] += $1; del[author] += $2 } /^$/ { author = ""; next } END { for (a in ins) { printf "%s,%d,%d,%d\n", a, commits[a], ins[a], del[a] } }'`)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalln("Error executing command:", err)
	}

	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")

	var authorStatsSlice []AuthorStats
	for _, line := range lines {
		data := strings.Split(line, ",")
		commitStats := AuthorStats{data[0], convertToInt(data[1]), convertToInt(data[2]), convertToInt(data[3])}
		authorStatsSlice = append(authorStatsSlice, commitStats)
	}

	return authorStatsSlice
}

func convertToInt(str string) int {
	value := 0
	_, err := fmt.Sscanf(str, "%d", &value)
	if err != nil {
		fmt.Println("Error converting to int:", err)
	}
	return value
}
