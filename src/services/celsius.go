package services

import (
	"fmt"

	"github.com/kentlouisetonino/tempscale/src/displays"
)

func Celsius() {
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
		fmt.Print(displays.Tab, "   Celsius Input (°C): ")
		_, err := fmt.Scan(&celsiusInput)

		if err != nil {
			celsiusInputError = true
			continue
		}
	}
}
