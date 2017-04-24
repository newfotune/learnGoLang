package tempconv

import "fmt"

//Celsius - Underlying type float64
type Celsius float64

//Fahrenheit - Underlying type float64
type Fahrenheit float64

const (
	//AbsoluteZeroC temp in Celcius
	AbsoluteZeroC Celsius = -273.15
	//FreezingC temp in Celcius
	FreezingC Celsius = 0
	//BoilingC temp in Celcuis
	BoilingC Celsius = 100
)

/**
 *Similar to toString() method in Java
 *A string representation of the Cencius type
 */
func (c Celsius) String() string {
	return fmt.Sprintf("%g*C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g*F", f)
}
