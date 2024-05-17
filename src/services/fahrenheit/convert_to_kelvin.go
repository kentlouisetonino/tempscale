package fahrenheit

func convertToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*(5.0/9.0) + 273.15
}
