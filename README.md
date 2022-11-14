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
	"github.com/mitchelllisle/redacted/redacted"
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

```
