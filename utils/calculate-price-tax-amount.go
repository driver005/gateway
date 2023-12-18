package utils

func CalculatePriceTaxAmount(price float64, taxRate float64, includesTax bool) float64 {
	if includesTax {
		return (price * taxRate) / (1 + taxRate)
	}

	return price * taxRate
}
