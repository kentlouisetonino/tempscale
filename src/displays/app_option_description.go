package displays

import "fmt"

func CelsiusOptionDescription() {
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
	AddNewline()
	fmt.Println(Tab, Tab, "Celsius to Fahrenheit - Kelvin - Rankine")
	AddNewline()
	fmt.Println(Tab, ColorBlue, "---------------------------------------------------------", ColorReset)
}
