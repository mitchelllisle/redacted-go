package redacted

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

func TestInfoTypes(t *testing.T) {
	simulations := 1000

	t.Run("match emails for Email InfoType", func(t *testing.T) {
		infoTypes := []InfoType{Email()}
		anonymiser := NewAnonymiser(infoTypes)

		for i := 1; i <= simulations; i++ {
			email := gofakeit.Email()
			matches := anonymiser.GetMatches(email)
			if len(matches) != 1 {
				t.Errorf("expected 1 matche; got %v", len(matches))
			}
		}
	})

	t.Run("match ids for AusDriversLicence InfoType", func(t *testing.T) {
		inf := AusDriversLicence()
		infoTypes := []InfoType{inf}
		anonymiser := NewAnonymiser(infoTypes)

		for i := 1; i <= simulations; i++ {
			generated := gofakeit.Regex(inf.Expr)
			matches := anonymiser.GetMatches(generated)
			if len(matches) != 1 {
				t.Errorf("expected 1 matche; got %v", len(matches))
			}
		}
	})
}
