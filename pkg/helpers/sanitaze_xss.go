package helpers

import (
	"html"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizeXss(s string) string {
	p := bluemonday.UGCPolicy()
	s = html.UnescapeString(s)
	description := p.Sanitize(s)
	return description
}
func SanitizeCustomeXss(s string) string {
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	s = html.UnescapeString(s)
	p.AllowElements("h1", "h2", "p", "strong", "em", "u", "s", "blockquote")
	description := p.Sanitize(s)
	return description
}
