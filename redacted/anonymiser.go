package redacted

import "regexp"

type Position struct {
	Start int
	End   int
}

type Match struct {
	Text      string
	Positions []Position
	InfoType  InfoType
}

type Anonymised struct {
	OriginalText   string
	AnonymisedText string
	Matches        []Match
}

type Anonymiser struct {
	Matchers []Matcher
}

func (a *Anonymiser) GetMatches(text string) []Match {
	var matches []Match

	for _, matcher := range a.Matchers {
		switch matcher.(type) {
		case *FuzzyMatcher:
			for _, m := range matcher.Match(text) {
				matches = append(matches, m)
			}
		case *RegexMatcher:
			for _, m := range matcher.Match(text) {
				matches = append(matches, m)
			}
		default:
			panic("Matcher must be either FuzzyMatcher or RegexMatcher")
		}
	}
	return matches
}

func (a *Anonymiser) Anonymise(text string) Anonymised {
	matches := a.GetMatches(text)
	newText := text
	for _, match := range matches {
		expr := regexp.MustCompile(match.Text)
		newText = expr.ReplaceAllLiteralString(newText, match.InfoType.Generate())
	}
	return Anonymised{
		OriginalText:   text,
		AnonymisedText: newText,
		Matches:        matches,
	}
}
