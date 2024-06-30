package card

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Duvewo/banquend/models"
)

// Generate generates a card
func Generate(
	IssuerIndustry int8,
	IssuerIDNumber int32,
	AccountNumber int64,
) (models.CardModel, error) {
	cardNumber, err := strconv.Atoi(
		fmt.Sprintf("%d%d%d", IssuerIndustry, IssuerIDNumber, AccountNumber),
	)

	if err != nil {
		return models.CardModel{}, err
	}

	cardNumberControlSum, err := strconv.Atoi(
		fmt.Sprintf("%d%d", cardNumber/10, checksum(cardNumber)),
	)

	if err != nil {
		return models.CardModel{}, err
	}

	return models.CardModel{
		PublicNumber: cardNumberControlSum,
	}, nil
}

// isValid check if n is valid according to the Luhn's algorithm
func checksum(n int) int {
	m, d, double, sum := 1.0, 10.0, 0.0, 0.0
	for i := 0; i < int(math.Log10(float64(n))+1); i++ {
		digit := math.Floor(math.Mod(float64(n), d) / m)

		if i%2 == 0 {
			sum += digit
			goto Next
		}

		double = digit * 2
		if double > 9 {
			sum += double - 9
			goto Next
		}

		sum += double

	Next:
		m *= 10
		d *= 10
	}
	return int(math.Mod(sum, 10))
}
