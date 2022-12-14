##  redacted

> 1️⃣ version: 0.2.0

> ✍️ author: Mitchell Lisle

📛 An experimental data anonymisation library

## Install


## Usage

### Anonymiser
```go
package main

import (
	"fmt"
	"github.com/mitchelllisle/redacted-go/redacted"
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
	// {{LAST_NAME}} {{LAST_NAME}} wendyschneider@tillman.com 2M46-yiS23441
}


```
