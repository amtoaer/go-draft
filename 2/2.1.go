package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

const (
	absoluteZeroC celsius = -273.15
	freezingC     celsius = 0
	boilingC      celsius = 100
)

func (c celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func cToF(c celsius) fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

func fToC(f fahrenheit) celsius {
	return celsius((f - 32) * 5 / 9)
}

func kToC(k kelvin) celsius {
	return celsius(k) + absoluteZeroC
}
func cToK(c celsius) kelvin {
	return kelvin(c - absoluteZeroC)
}

func main() {
	var ssd celsius
	var hsd fahrenheit
	var kel kelvin
	ssd = 100
	hsd = 50
	kel = 1
	fmt.Println(ssd, hsd, kel)
	fmt.Println(cToF(ssd))
	fmt.Println(cToK(ssd))
	fmt.Println(fToC(hsd))
	fmt.Println(kToC(kel))
}
