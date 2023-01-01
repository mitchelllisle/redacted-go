package redacted

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"testing"
)

func Range(min, max int) []int {
	out := make([]int, max-min)
	for i := range out {
		out[i] = min + i
	}
	return out
}

func RandBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

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

	t.Run("check postcode Regex matcher", func(t *testing.T) {
		infT := AusPostCode()
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		nswPostCodes := Range(2915, 4385)

		for _, postcode := range nswPostCodes {
			matches := anonymiser.GetMatches(fmt.Sprintf("%v", postcode))

			if len(matches) != 1 {
				t.Errorf(fmt.Sprintf("%v should have been matched by AusPostCode regex matcher", postcode))
			}
		}
	})

	t.Run("check integer length of LongDigit", func(t *testing.T) {
		infT := LongDigit(6)
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		for i := 1; i <= simulations; i++ {
			generated := RandBetween(100000, 9999999)
			matches := anonymiser.GetMatches(fmt.Sprintf("%v", generated))
			if len(matches) != 1 {
				t.Errorf("expected 1 match; got %v; for val %v", len(matches), generated)
			}
		}
	})

	t.Run("check integer length of LongDigit for vals less than max length", func(t *testing.T) {
		infT := LongDigit(3)
		matcher := NewRegexMatcher([]InfoType{infT})
		anonymiser := Anonymiser{Matchers: []Matcher{matcher}}

		for i := 1; i <= simulations; i++ {
			generated := RandBetween(0, 99)
			matches := anonymiser.GetMatches(fmt.Sprintf("%v", generated))
			if len(matches) != 0 {
				t.Errorf("expected 0 matches; got %v; for val %v", len(matches), generated)
			}
		}
	})
}
