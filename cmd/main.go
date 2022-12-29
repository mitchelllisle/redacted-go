package main

import (
	"fmt"
	"github.com/mitchelllisle/redacted/redacted"
)

func main() {
	text := "Millhouse Van Houten millhouse@milpool.com 2203 18423441"

	fuzzMatcher := redacted.NewFuzzyMatcher([]redacted.Phrase{
		{"Millhouse", "{{FIRST_NAME}}", 100},
		{"Van Houten", "{{LAST_NAME}}", 100},
	})
	regexMatcher := redacted.NewRegexMatcher([]redacted.InfoType{redacted.Email(), redacted.AusLicensePlate()})

	anonymiser := redacted.Anonymiser{Matchers: []redacted.Matcher{fuzzMatcher, regexMatcher}}

	out := anonymiser.Anonymise(text)

	fmt.Println(out.AnonymisedText)
}
