package redacted

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

func FakeInfoType() InfoType {
	name := "FakeInfoType"
	expr := "hello"
	return InfoType{
		Expr:         expr,
		Name:         name,
		WordBoundary: false,
		Generate:     func() string { return gofakeit.Regex(expr) },
	}
}

func TestHelloAnonymiser(t *testing.T) {
	infT := []InfoType{FakeInfoType()}
	matcher := NewRegexMatcher(infT)
	anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

	t.Run("one match in string", func(t *testing.T) {
		output := anonymiser.GetMatches("hello world")

		if len(output) != 1 {
			t.Errorf("expected 1 matche; got %v", len(output))
		}
	})

	t.Run("one match in string", func(t *testing.T) {
		output := anonymiser.GetMatches("hello hello world")

		if len(output) != 1 {
			t.Errorf("expected 1 matche; got %v", len(output))
		}
	})
}

func TestNewAnonymiser(t *testing.T) {
	infoTypes := []InfoType{Email(), AusDriversLicence()}
	matcher := NewRegexMatcher(infoTypes)
	anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

	t.Run("example string", func(t *testing.T) {
		text := "hello 000000 mitchell@lisle.com 000000 mitch@lisle.com"
		output := anonymiser.Anonymise(text)

		if output.OriginalText != text {
			t.Errorf("original text has been altered from: %s, to: %s", text, output.OriginalText)
		}

		if output.AnonymisedText == output.OriginalText {
			t.Errorf("anonymised text has not been altered altered original: %s, anon: %s", text, output.AnonymisedText)
		}

	})
}
