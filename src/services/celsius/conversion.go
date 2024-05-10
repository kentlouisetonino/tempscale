package celsius

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func Conversion() {
	var celsiusInput float64
	var celsiusInputError bool = false
	var conversionMenuInput int
	var backToMainMenu bool = false

	for backToMainMenu == false {
		// Clear the screen.
		displays.ClearScreen()
		displays.AddNewline()
		displays.AddNewline()

		// Show the description.
		displays.CelsiusOptionDescription()
		displays.AddNewline()

		// Error instruction.
		if celsiusInputError {
			displays.AppOptionInputError()
			displays.AddNewline()
		}

		// Ask the celsius input.
		fmt.Print(displays.Tab, "   ", "Celsius (°C)", displays.Tab, displays.Tab, ": ")
		_, err := fmt.Scan(&celsiusInput)

		if err != nil {
			celsiusInputError = true
			continue
		}

		// Display the equivalent of Celsius to Fahrenheit, Kelvin, Rankine.
		fmt.Print(displays.Tab, "   ", "Fahrenheit (°F)", displays.Tab, ": ", convertToFahrenheit(celsiusInput));
		displays.AddNewline()
		fmt.Print(displays.Tab, "   ", "Kelvin (°K)", displays.Tab, displays.Tab, ": ", convertToKelvin(celsiusInput));
		displays.AddNewline()
		fmt.Print(displays.Tab, "   ", "Rankine (°R)", displays.Tab, displays.Tab, ": ", convertToRankine(celsiusInput))
		displays.AddNewline()
		displays.AddNewline()

		// Ask if user wants to retry or go back to main menu.
		fmt.Print(displays.Tab, "   ", displays.ColorBlue, "[1-Retry, 2-Back, Any-Exit]", displays.ColorReset, displays.Tab, ": ")
		_, err = fmt.Scan(&conversionMenuInput)

		if err != nil {
			displays.ClearScreen()
			break
		}

		if conversionMenuInput == 1 {
			continue
		} else if conversionMenuInput == 2 {
			backToMainMenu = true
		} else {
			displays.ClearScreen()
			displays.AppExit()
		}
	}
}
