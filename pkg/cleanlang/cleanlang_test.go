package cleanlang

import "testing"

func TestValidator(t *testing.T) {
	validator := NewValidator()
	word := "l2o7"
	clean := validator.IsClean(word)
	if !clean {
		t.Errorf("expected '%s' to be clean", word)
	}
	word = "@ss"
	clean = validator.IsClean(word)
	if clean {
		t.Errorf("expected '%s' to not be clean", word)
	}
}
