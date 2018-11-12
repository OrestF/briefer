package analyzers

import (
	"strings"

	"github.com/urandom/text-summary/summarize"
)

type TextAnalyzer struct {
}

func (ta *TextAnalyzer) Brief(title string, text string) string {
	s := summarize.NewFromString(title, text)

	return strings.Join(s.KeyPoints(), " ")
}
