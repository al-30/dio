package main

import "fmt"

func convertKelvinToCelsius(k float32) float32 {
	return (k - 273.0)
}
func main() {
	c := convertKelvinToCelsius(373)
	fmt.Println(c)

}
