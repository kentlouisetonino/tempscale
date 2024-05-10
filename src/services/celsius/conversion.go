package celsius

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func Conversion() {
	var celsiusInput float64
	var celsiusInputError bool = false

	for {
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
		break
	}
}
