package encoder

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
