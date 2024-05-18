package kelvin

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func Conversion() {
	var kelvinInput float64
	var celsiusInputError bool = false
	var conversionMenuInput int
	var backToMainMenu bool = false

	for backToMainMenu == false {
		// Clear the screen.
		displays.ClearScreen()
		displays.AddNewline()
		displays.AddNewline()

		// Show the description.
		displays.FahrenheitOptionDescription()
		displays.AddNewline()

		// Error instruction.
		if celsiusInputError {
			displays.AppOptionInputError()
			displays.AddNewline()
		}

		// Ask the celsius input.
		fmt.Print(displays.Tab, "   ", "Kelvin (°F)", displays.Tab, displays.Tab, ": ")
		_, err := fmt.Scan(&kelvinInput)

		if err != nil {
			celsiusInputError = true
			continue
		}

		// Display the equivalent of Celsius to Fahrenheit, Kelvin, Rankine.
		fmt.Print(displays.Tab, "   ", "Celsius (°C)", displays.Tab, displays.Tab, ": ", displays.ColorGreen, convertToCelsius(kelvinInput), displays.ColorReset);
		displays.AddNewline()
		fmt.Print(displays.Tab, "   ", "Fahrenheit (°F)", displays.Tab, ": ", displays.ColorGreen, convertToFahrenheit(kelvinInput), displays.ColorReset);
		displays.AddNewline()
		fmt.Print(displays.Tab, "   ", "Rankine (°R)", displays.Tab, displays.Tab, ": ", displays.ColorGreen, convertToRankine(kelvinInput), displays.ColorReset)
		displays.AddNewline()
		displays.AddNewline()

		// Ask if user wants to retry or go back to main menu.
		fmt.Print(displays.Tab, "   ", displays.ColorBlue, "[1-Retry, 2-Back, Any-Exit]", displays.ColorReset, displays.Tab, ": ")
		_, err = fmt.Scan(&conversionMenuInput)

		if err != nil {
			displays.ClearScreen()
			displays.AppExit()
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
