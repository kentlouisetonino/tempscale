package main

import (
	"github.com/kentlouisetonino/tempscale/src/displays"
	"github.com/kentlouisetonino/tempscale/src/services/celsius"
	"github.com/kentlouisetonino/tempscale/src/services/fahrenheit"
	"github.com/kentlouisetonino/tempscale/src/services/kelvin"
	"github.com/kentlouisetonino/tempscale/src/services/rankine"
)

func main() {
	optionInput := ""
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
			displays.AppOptionInputError()
			displays.AddNewline()
		}

		// Display the options.
		displays.AppOptions()
		displays.AddNewline()

		// Choose a temperature scale to convert.
		optionInput = displays.AppOptionInput()
		displays.AddNewline()
		displays.AddNewline()
		displays.AddNewline()

		if optionInput == "-1" {
			invalidInput = true
			continue
		}

		if optionInput == "1" {
			celsius.Conversion()
		}

		if optionInput == "2" {
			fahrenheit.Conversion()
		}

		if optionInput == "3" {
			kelvin.Conversion()
		}

		if optionInput == "4" {
			rankine.Conversion()
		}

		if optionInput == "4" {
			displays.ClearScreen()
			displays.AppExit()
		}
	}
}
