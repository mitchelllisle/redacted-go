package redacted

import (
	"testing"
)

func TestNewFuzzyMatcher(t *testing.T) {
	validMatcher := NewFuzzyMatcher([]Phrase{{"hello", "{GREETING}", 100}})
	//thresholdMatcher := NewFuzzyMatcher([]Phrase{{"missing", "{MISSING}", 1}})

	t.Run("test new fuzzy matches", func(t *testing.T) {
		matched := validMatcher.Match("hello")
		if len(matched) != 1 {
			t.Errorf("expected 1 match; got %v", len(matched))
		}

		if matched[0].Text != "hello" {
			t.Errorf("expected match test to be 'hello'; not %v", matched[0].Text)
		}

		generated := matched[0].InfoType.Generate()
		if generated != "{GREETING}" {
			t.Errorf("expected {GREETING} as generated text; not %v", generated)
		}
	})

}
