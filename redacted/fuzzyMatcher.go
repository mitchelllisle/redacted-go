package redacted

import (
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"regexp"
)

type Phrase struct {
	Phrase      string
	Replacement string
	Threshold   int
}

type FuzzyMatchPhrase struct {
	Phrase    string
	Threshold int
	Generate  func() string
}

type FuzzyMatcher struct {
	Phrases []FuzzyMatchPhrase
}

func makeFuzzyInfoType(text string, generate func() string) InfoType {
	return InfoType{
		Expr:         text,
		Name:         "FuzzyMatch",
		WordBoundary: false,
		Generate:     generate,
	}
}

func (m *FuzzyMatcher) Match(text string) []Match {
	var matches []Match
	for _, p := range m.Phrases {
		fuzzyInfoType := makeFuzzyInfoType(p.Phrase, p.Generate)

		found := fuzzy.RankFind(p.Phrase, []string{text})
		for _, match := range found {
			if match.Distance <= p.Threshold {
				expr, err := regexp.Compile(p.Phrase)
				PanicOnError(err, fmt.Sprintf("unable to create regex expression from %s", p.Phrase))
				positions := getPositionsForExpr(expr, text)
				m := Match{
					Text:      match.Source,
					Positions: positions,
					InfoType:  fuzzyInfoType,
				}
				matches = append(matches, m)
			}
		}
	}
	return matches
}

func NewFuzzyMatcher(phrases []Phrase) *FuzzyMatcher {
	var fuzzyPhrases []FuzzyMatchPhrase
	for _, p := range phrases {
		replacement := p.Replacement
		genFunc := func() string { return replacement }

		fuzzyPhrases = append(fuzzyPhrases, FuzzyMatchPhrase{
			Phrase:    p.Phrase,
			Generate:  genFunc,
			Threshold: p.Threshold,
		})
	}
	return &FuzzyMatcher{Phrases: fuzzyPhrases}
}
