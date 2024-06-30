package card_test

import (
	"fmt"
	"testing"

	"github.com/Duvewo/banquend/internal/card"
)

func TestGenerate(t *testing.T) {
	fmt.Println(card.Generate(4, 32168, 4123536234))
}
