package generate

import (
	"strings"
	"testing"
)

func TestDefaultGenerator(t *testing.T) {

	g := Generator{
		Alpha:     true,
		Numeric:   true,
		Symb:      true,
		UpperOnly: false,
		LowerOnly: false,
	}
	t.Run("generate password longer than 12 characters", func(t *testing.T) {
		got := g.Generate()
		if len(got) < 12 {
			t.Errorf("password length not greater than 12")
		}
	})
	t.Run("generate different passwords", func(t *testing.T) {
		got1 := g.Generate()
		got2 := g.Generate()

		if got1 == got2 {
			t.Errorf("passwords should be different, got %q", got1)
		}
	})

	t.Run("password should contain alphabetical characters, numbers, and symbols", func(t *testing.T) {
		got := g.Generate()

		categories := make(map[string]int)
		categories["alphabet"] = 0
		categories["number"] = 0
		categories["symbol"] = 0

		for _, val := range got {

			if strings.ContainsRune(ALPHA, val) {
				categories["alphabet"] += 1
			}
			if strings.ContainsRune(NUM, val) {
				categories["number"] += 1
			}
			if strings.ContainsRune(SYMB, val) {
				categories["symbol"] += 1
			}
		}

		if categories["alphabet"] == 0 {
			t.Errorf("no alphabetical characters found in password")
		}
		if categories["number"] == 0 {
			t.Errorf("no numbers found in password")
		}
		if categories["symbol"] == 0 {
			t.Errorf("no symbols found in password")
		}

	})

}
