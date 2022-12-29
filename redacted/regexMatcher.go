package redacted

import (
	"fmt"
	"regexp"
	"strings"
)

type RegexMatcher struct {
	Expr          *regexp.Regexp
	InfoTypes     []InfoType
	infoTypeMap   map[string]InfoType
	infoTypeNames []string
	Name          string
	WordBoundary  bool
	Generate      func() string
}

func getPositionsForExpr(expr *regexp.Regexp, text string) []Position {
	var positions []Position
	for _, index := range expr.FindAllStringIndex(text, -1) {
		positions = append(positions, Position{
			Start: index[0],
			End:   index[1],
		})
	}
	return positions
}

func createSubExpression(infoType InfoType, i int) strings.Builder {
	var builder strings.Builder
	if i > 0 {
		builder.WriteString(`|`)
	}
	_, err := builder.WriteString(`(?P<`)

	// write Expr name
	_, err = builder.WriteString(fmt.Sprintf(`%s>`, infoType.Name))

	if infoType.WordBoundary {
		_, err = builder.WriteString(`\b`)
	}
	_, err = builder.WriteString(infoType.Expr)
	_, err = builder.WriteString(`)`)
	PanicOnError(err, fmt.Sprintf("unable to create sub expression for %s", infoType.Name))
	return builder
}

func (m *RegexMatcher) getUniqueMatches(text string) []string {
	var uniqueMatches = make(map[string]bool)
	var keys []string
	for _, match := range m.Expr.FindAllString(text, -1) {
		uniqueMatches[match] = true
	}
	for k := range uniqueMatches {
		keys = append(keys, k)
	}
	return keys
}

func (m *RegexMatcher) getGroupName(text string) (bool, string) {
	for _, match := range m.Expr.FindAllStringSubmatch(text, -1) {
		for groupIdx, group := range match {
			name := m.infoTypeNames[groupIdx]
			if (name != "") && (group != "") {
				return true, name
			}
		}
	}
	return false, ""
}

func (m *RegexMatcher) Match(text string) []Match {
	var matches []Match
	matchesAsString := m.getUniqueMatches(text)

	for _, match := range matchesAsString {
		expr := regexp.MustCompile(match)
		ok, name := m.getGroupName(match)
		if ok {
			positions := getPositionsForExpr(expr, text)
			matches = append(matches, Match{
				Text:      match,
				Positions: positions,
				InfoType:  m.infoTypeMap[name],
			})
		}
	}
	return matches
}

func NewRegexMatcher(infoTypes []InfoType) *RegexMatcher {
	var builder strings.Builder

	for i, inf := range infoTypes {
		subExpr := createSubExpression(inf, i)
		builder.WriteString(subExpr.String())
	}
	r, err := regexp.Compile(builder.String())
	PanicOnError(err, "unable to add create regex expression from string")

	infoTypeMap := make(map[string]InfoType)
	for _, infT := range infoTypes {
		infoTypeMap[infT.Name] = infT
	}
	names := r.SubexpNames()
	return &RegexMatcher{
		Expr:          r,
		InfoTypes:     infoTypes,
		infoTypeMap:   infoTypeMap,
		infoTypeNames: names,
		WordBoundary:  false,
		Generate:      func() string { return "" },
	}
}
