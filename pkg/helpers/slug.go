package helpers

import "strings"

func Slug(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
