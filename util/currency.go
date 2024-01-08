package util

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

func ValidateCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}

	return false
}
