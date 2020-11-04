package cleanlang

type Validator interface {
	IsClean(word string) bool
}

type ValidatorImpl struct {
	badWords map[string]bool
}

type Words struct {
	Words []string `json:"words"`
}

func NewValidator() Validator {
	return &ValidatorImpl{badWords: words}
}

func (v *ValidatorImpl) IsClean(word string) bool {
	_, found := v.badWords[word]
	return !found
}
