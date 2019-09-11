package Testing

func calculateTax(value int) float64 {
	var tax float64
	switch true {
	case value <= 12500:
		tax = 0
	case 12500 < value && value <= 50000:
		tax = float64(value - 12500) * 0.2
	case 50000 < value && value <= 150000:
		tax = 7500 + float64(value - 50000) * 0.4
	case value > 150000:
		tax = 47500 + float64(value - 150000) * 0.45
	}
	return tax
}