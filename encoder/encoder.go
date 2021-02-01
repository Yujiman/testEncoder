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

func (e *Encoder) Encrypt(text string, key int) (string, error) {
	result := ""
	for _, char := range text {
		result += string(e.getEncryptChar(char, key))
	}
	return result, nil
}
func (e *Encoder) getEncryptChar(char rune, key int) rune {

	if unicode.IsLower(char) {
		idx, err := e.getIdxLowerList(char)
		if err != nil {
			return char
		}
		shift := key % len(e.lowerList)

		char := e.lowerList[(idx+shift)%len(e.lowerList)]
		return char
	}

	if unicode.IsUpper(char) {
		idx, err := e.getIdxCapitalList(char)
		if err != nil {
			return char
		}

		shift := key % len(e.capitalList)

		char = e.capitalList[(idx+shift)%len(e.capitalList)]

		return char
	}
	return char
}

func (e *Encoder) Decrypt(text string, key int) (string, error) {
	result := ""
	for _, char := range text {
		result += string(e.getDecryptChar(char, key))
	}
	return result, nil
}
func (e *Encoder) getDecryptChar(char rune, key int) rune {
	if unicode.IsLower(char) {
		idx, err := e.getIdxLowerList(char)
		if err != nil {
			return char
		}
		shift := key % len(e.lowerList)

		char := e.lowerList[(idx-shift)%len(e.lowerList)]
		return char
	}

	if unicode.IsUpper(char) {
		idx, err := e.getIdxCapitalList(char)
		if err != nil {
			return char
		}

		shift := key % len(e.capitalList)

		char = e.capitalList[(idx-shift)%len(e.capitalList)]

		return char
	}
	return char
}

func (e *Encoder) getIdxLowerList(char rune) (int, error) {
	for index, c := range e.lowerList {
		if char == c {
			return index, nil
		}
	}
	return 0, errors.New("character won't find in")
}

func (e *Encoder) getIdxCapitalList(char rune) (int, error) {
	for index, c := range e.capitalList {
		if char == c {
			return index, nil
		}
	}
	return 0, errors.New("character won't find in ")
}
