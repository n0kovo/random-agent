package main

import (
	"bufio"
	"embed"
	"fmt"
	"math/rand"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
)

//go:embed user-agents.txt
var userAgentsFS embed.FS

var (
	userAgentsList []string
)

func main() {
	var matchStrings, filterStrings []string
	flag.StringSliceVarP(&matchStrings, "match", "m", nil, "Only return user-agents matching the specified string. Can be specified multiple times.")
	flag.StringSliceVarP(&filterStrings, "filter", "f", nil, "Filter out user-agents matching the specified string. Can be specified multiple times.")
	flag.Parse()

	userAgentsList, _ = loadUserAgents()

	random_source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(random_source)

	rand.Shuffle(len(userAgentsList), func(i, j int) {
		userAgentsList[i], userAgentsList[j] = userAgentsList[j], userAgentsList[i]
	})

	for _, userAgent := range userAgentsList {
		if len(matchStrings) == 0 && len(filterStrings) == 0 {
			fmt.Print(userAgent)
			break
		} else if matchesFilter(userAgent, matchStrings) && !matchesFilter(userAgent, filterStrings) {
			fmt.Print(userAgent)
			break
		}
	}
}

func loadUserAgents() ([]string, error) {
	file, err := userAgentsFS.Open("user-agents.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func matchesFilter(userAgent string, filters []string) bool {
	for _, filter := range filters {
		if strings.Contains(userAgent, filter) {
			return true
		}
	}
	return len(filters) == 0
}
