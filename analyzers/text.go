package analyzers

import (
	"github.com/urandom/text-summary/summarize"
)

type TextAnalyzer struct {
}

func (ta *TextAnalyzer) Brief(text string) ([]string){
	s := summarize.NewFromString("Title", text)

	println(s.KeyPoints())
	return s.KeyPoints()
}
