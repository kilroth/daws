package utils

import "strings"

func Slugify(name string) string {
	tmpName := strings.ReplaceAll(name, " ", "-")
	tmpName = strings.ToLower(tmpName)
	return tmpName
}
