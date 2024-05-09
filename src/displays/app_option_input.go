package displays

import "fmt"

func AppOptionInput() string {
	optionInput := ""

	fmt.Print(Tab, ColorBlue, "   Choose an option number: ", ColorReset)
	fmt.Scan(&optionInput)

	// Celsius option.
	if optionInput == "1" {
		return "1"
	}

	// Fahrenheit option.
	if optionInput == "2" {
		return "2"
	}

	// Kelvin option.
	if optionInput == "3" {
		return "3"
	}

	// Rankin option.
	if optionInput == "4" {
		return "4"
	}

	// Exit application option.
	if optionInput == "5" {
		return "5"
	}

	// Invalid input.
	return "-1"
}
