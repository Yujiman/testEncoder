package encoder

import (
	"errors"
	"unicode"
)

type Encoder struct {
	lowerList, capitalList []rune
}

func NewEncoder(lowList, capList []rune) *Encoder {
	return &Encoder{
		lowerList:   lowList,
		capitalList: capList,
	}
}

func (cl *Encoder) Encrypt(text string, key int) (string, error) {
	return "", nil
}

func (cl *Encoder) getNewChar(char rune, shift int) rune {
	if unicode.IsLower(char) {
		return char
	}

	if unicode.IsUpper(char) {
		return char
	}
	return char
}
func (cl *Encoder) getIdxLowerList(char rune) (int, error) {
	for index, c := range cl.lowerList {
		if char == c {
			return index, nil
		}
	}
	return 0, errors.New("character won't find in")
}

func (cl *Encoder) getIdxCapitalList(char rune) (int, error) {
	for index, c := range cl.capitalList {
		if char == c {
			return index, nil
		}
	}
	return 0, errors.New("character won't find in ")
}
