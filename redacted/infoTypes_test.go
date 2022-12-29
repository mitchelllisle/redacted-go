package redacted

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

func TestInfoTypes(t *testing.T) {
	simulations := 1000

	t.Run("match emails for Email InfoType", func(t *testing.T) {
		infT := Email()
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		for i := 1; i <= simulations; i++ {
			email := gofakeit.Email()
			matches := anonymiser.GetMatches(email)
			if len(matches) != 1 {
				t.Errorf("expected 1 matche; got %v", len(matches))
			}
		}
	})

	t.Run("match ids for AusDriversLicence InfoType", func(t *testing.T) {
		infT := AusDriversLicence()
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		for i := 1; i <= simulations; i++ {
			generated := gofakeit.Regex(infT.Expr)
			matches := anonymiser.GetMatches(generated)
			if len(matches) != 1 {
				t.Errorf("expected 1 matche; got %v", len(matches))
			}
		}
	})

	t.Run("match ids for AusPassport InfoType", func(t *testing.T) {
		infT := AusPassport()
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		for i := 1; i <= simulations; i++ {
			generated := gofakeit.Regex(infT.Expr)
			matches := anonymiser.GetMatches(generated)
			if len(matches) != 1 {
				t.Errorf("expected 1 matche; got %v", len(matches))
			}
		}
	})

	t.Run("match ids for AusLicensePlate InfoType", func(t *testing.T) {
		infT := AusLicensePlate()
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		for i := 1; i <= simulations; i++ {
			generated := gofakeit.Regex(infT.Expr)
			matches := anonymiser.GetMatches(generated)
			if len(matches) != 1 {
				t.Errorf("expected 1 matche; got %v", len(matches))
			}
		}
	})

	t.Run("match ids for AusTaxFileNumber InfoType", func(t *testing.T) {
		infT := AusTaxFileNumber()
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		for i := 1; i <= simulations; i++ {
			generated := gofakeit.Regex(infT.Expr)
			matches := anonymiser.GetMatches(generated)
			if len(matches) != 1 {
				t.Errorf("expected 1 matche; got %v", len(matches))
			}
		}
	})
}
