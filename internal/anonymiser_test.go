package internal

import (
	"regexp"
	"testing"
)

func TestHelloAnonymiser(t *testing.T) {
	r, _ := regexp.Compile("hello")
	anonymiser := Anonymiser{Expr: r}

	t.Run("one match in string", func(t *testing.T) {
		output := anonymiser.GetMatches("hello world")

		if len(output) != 1 {
			t.Errorf("expected 1 matche; got %v", len(output))
		}
	})

	t.Run("one match in string", func(t *testing.T) {
		output := anonymiser.GetMatches("hello hello world")

		if len(output) != 2 {
			t.Errorf("expected 2 matche; got %v", len(output))
		}
	})

}
