package displays

import "fmt"

func AppMainOption() string {
	option := ""

	fmt.Print(Tab, ColorBlue, "   Choose an option number: ", ColorReset)
	fmt.Scan(&option)

	// Celsius option.
	if option == "1" {
		return "1"
	}

	// Fahrenheit option.
	if option == "2" {
		return "2"
	}

	// Kelvin option.
	if option == "3" {
		return "3"
	}

	// Rankin option.
	if option == "4" {
		return "4"
	}

	// Exit application option.
	if option == "5" {
		return "5"
	}

	// Invalid input.
	return "-1"
}
