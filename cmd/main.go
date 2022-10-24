package main

import (
	"redacted-go/internal"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("hello")

	anonymiser := internal.Anonymiser{Expr: r}

	text := "hello hello world"
	matches := anonymiser.GetMatches(text)

	println(matches)
}
