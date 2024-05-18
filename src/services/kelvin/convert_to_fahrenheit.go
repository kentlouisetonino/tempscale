package kelvin

func convertToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*(9.0/5.0) + 32
}
