package displays

import (
	"fmt"
	"strconv"
)

const openBracket = ColorBlue + "[" + ColorReset
const closeBracket = ColorBlue + "]" + ColorReset

func optionNumber(number int) string {
	return ColorGreen + strconv.Itoa(number) + ColorReset
}

func AppOptions() {
	fmt.Println(Tab, " ", openBracket, optionNumber(1), closeBracket, "Celsius (°C)")
	fmt.Println(Tab, " ", openBracket, optionNumber(2), closeBracket, "Fahrenheit (°F)")
	fmt.Println(Tab, " ", openBracket, optionNumber(3), closeBracket, "Kelvin (°K)")
	fmt.Println(Tab, " ", openBracket, optionNumber(4), closeBracket, "Rankine (°R)")
	fmt.Println(Tab, " ", openBracket, optionNumber(5), closeBracket, "Exit Application")
}
