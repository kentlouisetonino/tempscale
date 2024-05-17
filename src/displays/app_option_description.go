package displays

import "fmt"

func CelsiusOptionDescription() {
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
	AddNewline()
	AddNewline()
	fmt.Println(Tab, Tab, ColorGreen, "Celsius", ColorReset, "to Fahrenheit - Kelvin - Rankine")
	AddNewline()
	AddNewline()
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
}

func FahrenheitOptionDescription() {
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
	AddNewline()
	AddNewline()
	fmt.Println(Tab, Tab, ColorGreen, "Celsius", ColorReset, "to Fahrenheit - Kelvin - Rankine")
	AddNewline()
	AddNewline()
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
}
