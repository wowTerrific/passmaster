package generate

import (
	"math/rand"
	"strings"
)

const (
	ALPHA string = "abcdefghijklmnopqrstuvwxyz"
	NUM   string = "1234567890"
	SYMB  string = "!?()=+"
)

func Generate() string {

	min := 12
	max := 16
	wordLimit := rand.Intn(max-min+1) + min

	var password strings.Builder

	for i := 0; i < wordLimit; i++ {
		password.WriteByte(ALPHA[i])
	}

	return password.String()
}
