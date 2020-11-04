package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const wordsFile = "words.json"

type Validator interface {
	IsClean(word string) bool
}

type ValidatorImpl struct {
	badWords map[string]bool
}

type Words struct {
	Words []string `json:"words"`
}

func NewValidator() (Validator, error) {
	jsonFile, err := os.Open(wordsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s'", wordsFile)
	}
  byteValue, err := ioutil.ReadAll(jsonFile)
  var wordIndex Words
	err = json.Unmarshal(byteValue, &wordIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarashal words in '%s'", wordsFile)
	}
	badWords := map[string]bool{}
	for _, word := range wordIndex.Words {
		badWords[word] = true
	}
	return &ValidatorImpl{badWords: badWords}, nil
}

func (v *ValidatorImpl) IsClean(word string) bool {
	_, found := v.badWords[word]
	return found
}
