package internal

import (
	"regexp"
)

type Anonymiser struct {
	Expr *regexp.Regexp
}

type Match struct {
	Text  string
	Start int
	End   int
}

func (a *Anonymiser) GetMatches(text string) []Match {
	var output []Match
	matches := a.Expr.FindAllStringIndex(text, -1)

	for _, match := range matches {
		start := match[0]
		end := match[1]
		found := text[match[0]:match[1]]
		output = append(output, Match{Text: found, Start: start, End: end})
	}
	return output
}
