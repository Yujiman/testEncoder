package encoder

import "unicode"

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
