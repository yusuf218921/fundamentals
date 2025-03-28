package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i<10000; {
		if z*z > x {
			z -= (z*z - x) / (2 * z)
		} else {
			z++
		}
		i++
	}
	return z
}

func main() {
	fmt.Println("2'nin kare kökü -> ",Sqrt(2))
}