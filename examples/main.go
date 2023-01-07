package main

import (
	"fmt"
	"github.com/mitchelllisle/redacted-go/redacted"
)

type Transaction struct {
	Description string
	Name        string
	Identifier  string
	Email       string
}

func main() {
	transactions := redacted.ReadCSV[Transaction]("examples/data.csv", true)

	for _, transaction := range transactions {
		fuzzMatcher := redacted.NewFuzzyMatcher([]redacted.Phrase{
			{transaction.Name, "{{NAME}}", 100},
			{transaction.Identifier, "{{IDENTIFIER}}", 100},
		})
		regexMatcher := redacted.NewRegexMatcher([]redacted.InfoType{redacted.Email(), redacted.AusDriversLicence()})

		anonymiser := redacted.Anonymiser{Matchers: []redacted.Matcher{fuzzMatcher, regexMatcher}}

		out := anonymiser.Anonymise(transaction.Description)

		fmt.Println(fmt.Sprintf("Found %v matches for the following:", len(out.Matches)))
		fmt.Println(fmt.Sprintf("Original: %s", transaction.Description))
		fmt.Println(fmt.Sprintf("Anonynised: %s", out.AnonymisedText))
		fmt.Println()
	}
}
