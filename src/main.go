package main

import (
	"github.com/kentlouisetonino/tempscale/src/displays"
	"github.com/kentlouisetonino/tempscale/src/services"
)

func main() {
	mainInput := ""
	invalidInput := false

	for {
		// Clear the terminal first.
		displays.ClearScreen()
		displays.AddNewline()
		displays.AddNewline()

		// Display the application description.
		displays.AppDescription()
		displays.AddNewline()

		// Error instruction.
		if invalidInput {
			displays.AppMainInputError()
			displays.AddNewline()
		}

		// Display the options.
		displays.AppOptions()
		displays.AddNewline()

		// Choose a temperature scale to convert.
		mainInput = displays.AppMainInput()
		displays.AddNewline()
		displays.AddNewline()
		displays.AddNewline()

		if mainInput == "-1" {
			invalidInput = true
			continue
		}

		if mainInput == "1" {
			services.Celsius()
			break
		}
	}
}
