package pkg

import "testing"

func TestValidator(t *testing.T) {
	validator, err := NewValidator()
	if err != nil {
		t.Errorf("failed to create validator")
	}
	word := "abc"
	clean := validator.IsClean(word)
	if clean {
		t.Errorf("expected '%s' to be clean", word)
	}
	word = "@ss"
	clean = validator.IsClean(word)
	if clean {
		t.Errorf("expected '%s' to not be clean", word)
	}
}
