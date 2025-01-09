package generate

import (
	"math/rand"
	"strings"
)

const (
	ALPHA string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NUM   string = "1234567890"
	SYMB  string = "!?()=+$"
)

type Generator struct {
	Alpha     bool
	Numeric   bool
	Symb      bool
	UpperOnly bool
	LowerOnly bool
}

func (g *Generator) Generate() string {

	min := 14
	max := 19 // non inclusive
	wordLimit := rand.Intn(max-min) + min
	categories := []string{ALPHA, NUM, SYMB}

	var password strings.Builder

	for i := 0; i < wordLimit; i++ {
		category := categories[rand.Intn(len(categories))]
		character := category[rand.Intn(len(category))]

		password.WriteByte(character)
	}

	return password.String()
}
