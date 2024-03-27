package displays

import "fmt"

func AppOptions() {
	fmt.Println(Tab, " ", "[1]", "Celsius (°C)")
	fmt.Println(Tab, " ", "[2]", "Fahrenheit (°F)")
	fmt.Println(Tab, " ", "[3]", "Kelvin (°K)")
	fmt.Println(Tab, " ", "[4]", "Rankine (°R)")
}
