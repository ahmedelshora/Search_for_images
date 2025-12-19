package main

import (
	"regexp"
	"strings"
)

func sanitizeFilename(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	re := regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	return re.ReplaceAllString(name, "")
}
