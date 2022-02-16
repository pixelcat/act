package cmd

import (
	"strings"
)

func (i *Input) newPlatforms() map[string]string {
	platforms := map[string]string{
		"ubuntu-latest":  "node:12-buster-slim",
		"ubuntu-20.04":   "node:12-buster-slim",
		"ubuntu-18.04":   "node:12-buster-slim",
		"ubuntu-16.04":   "node:12-stretch-slim",
		"windows-latest": "",
		"windows-2019":   "",
		"macos-latest":   "",
		"macos-10.15":    "",
	}

	for _, p := range i.platforms {
		pParts := strings.Split(p, "=")
		if len(pParts) == 2 {
			platforms[pParts[0]] = pParts[1]
		}
	}
	return platforms
}
