package redacted

type Matcher interface {
	Match(string) []Match
}
