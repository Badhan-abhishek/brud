package stringutils

import (
	"regexp"
	"strings"
)

var (
	nonAlphanumeric = regexp.MustCompile(`[^a-z0-9\-]+`)
	multiDash       = regexp.MustCompile(`-{2,}`)
	spaceToDash     = regexp.MustCompile(`[\s_]+`)
)

func Slugify(input string) string {
	slug := strings.ToLower(input)
	slug = spaceToDash.ReplaceAllString(slug, "-")
	slug = nonAlphanumeric.ReplaceAllString(slug, "")
	slug = multiDash.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	return slug
}

func Truncate(input string, maxLength int) string {
	if len(input) <= maxLength {
		return input
	}
	return input[:maxLength] + "..."
}
