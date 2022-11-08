package main

import (
	"fmt"
	"redacted-go/redacted"
)

func main() {
	infoTypes := []redacted.InfoType{
		redacted.AusDriversLicence(),
		redacted.Email(),
		redacted.AusTaxFileNumber(),
	}

	anonymiser := redacted.NewAnonymiser(infoTypes)

	text := "hello 000000 mitchell@lisle.com 000000 mitch@lisle.com"
	anonymised := anonymiser.Anonymise(text)
	fmt.Println(anonymised)
}
