package pkg

import (
	"fmt"
	"strings"
)

// ParseShortCode parses a short org/repo to a GitHub URL
func ParseShortCode(args []string, output func(string) string) []string {
	return ParseShortCodeWithProtocol(args, "https", output)
}

// ParseShortCodeWithProtocol parses a short org/repo with protocol
func ParseShortCodeWithProtocol(args []string, protocol string, output func(string) string) []string {
	if len(args) <= 0 {
		return args
	}
	if output == nil {
		output = defaultOutput
	}

	var result []string
	switch protocol {
	case "https":
		result = []string{fmt.Sprintf("https://github.com/%s", args[0])}
	case "ssh", "git":
		result = []string{fmt.Sprintf("git@github.com:%s.git", args[0])}
	}

	if len(args) > 1 {
		result = append(result, args[1:]...)
	} else {
		if target := output(args[0]); target != "" {
			result = append(result, target)
		}
	}
	return result
}

func defaultOutput(arg string) string {
	if orgAndRepo := strings.Split(arg, "/"); len(orgAndRepo) == 2 {
		return orgAndRepo[1]
	}
	return ""
}
