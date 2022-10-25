package main

import (
	"redacted-go/internal"
)

func main() {
	infoTypes := []internal.InfoType{
		internal.AusDriversLicence(),
		internal.Email(),
	}

	anonymiser := internal.NewAnonymiser(infoTypes)

	text := "hello 000000 mitchell@lisle.com"
	matches := anonymiser.GetMatches(text)

	for _, m := range matches {
		println(m.Text)
	}
}
