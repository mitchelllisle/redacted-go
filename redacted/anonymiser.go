package redacted

import (
	"fmt"
	"regexp"
	"strings"
)

type Anonymiser struct {
	Expr          *regexp.Regexp
	InfoTypes     []InfoType
	infoTypeMap   map[string]InfoType
	infoTypeNames []string
}

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

func NewAnonymiser(infoTypes []InfoType) *Anonymiser {
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
	return &Anonymiser{Expr: r, InfoTypes: infoTypes, infoTypeMap: infoTypeMap, infoTypeNames: names}
}

func (a *Anonymiser) getUniqueMatches(text string) []string {
	var uniqueMatches = make(map[string]bool)
	var keys []string
	for _, match := range a.Expr.FindAllString(text, -1) {
		uniqueMatches[match] = true
	}
	for k := range uniqueMatches {
		keys = append(keys, k)
	}
	return keys
}

func (a *Anonymiser) getGroupName(text string) (bool, string) {
	for _, match := range a.Expr.FindAllStringSubmatch(text, -1) {
		for groupIdx, group := range match {
			name := a.infoTypeNames[groupIdx]
			if (name != "") && (group != "") {
				return true, name
			}
		}
	}
	return false, ""
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

func (a *Anonymiser) GetMatches(text string) []Match {
	var matches []Match
	matchesAsString := a.getUniqueMatches(text)

	for _, match := range matchesAsString {
		expr := regexp.MustCompile(match)
		ok, name := a.getGroupName(match)
		if ok {
			positions := getPositionsForExpr(expr, text)
			matches = append(matches, Match{
				Text:      match,
				Positions: positions,
				InfoType:  a.infoTypeMap[name],
			})
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
