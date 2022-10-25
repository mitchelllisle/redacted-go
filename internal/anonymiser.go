package internal

import (
	"fmt"
	"regexp"
	"strings"
)

type Anonymiser struct {
	Expr      *regexp.Regexp
	InfoTypes []InfoType
}

type Match struct {
	Text  string
	Start int
	End   int
}

func createSubExpression(infoType InfoType, i int) strings.Builder {
	var builder strings.Builder
	if i > 0 {
		builder.WriteString(`|`)
	}
	_, err := builder.WriteString(`(?P<`)

	// write expr name
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
	return &Anonymiser{Expr: r, InfoTypes: infoTypes}
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
