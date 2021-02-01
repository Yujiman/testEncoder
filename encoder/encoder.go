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
	result := ""
	for _, char := range text {
		result += string(cl.getNewChar(char, key))
	}
	return result, nil
}

func (cl *Encoder) getNewChar(char rune, key int) rune {

	if unicode.IsLower(char) {
		idx, err := cl.getIdxLowerList(char)
		if err != nil {
			return char
		}
		shift := key % len(cl.lowerList)

		char := cl.lowerList[(idx+shift)%len(cl.lowerList)]
		return char
	}

	if unicode.IsUpper(char) {
		idx, err := cl.getIdxCapitalList(char)
		if err != nil {
			return char
		}

		shift := key % len(cl.capitalList)

		char = cl.capitalList[(idx+shift)%len(cl.capitalList)]

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
