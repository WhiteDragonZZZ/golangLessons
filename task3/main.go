package main

import (
	"fmt"
)

func main() {
	var valueVelocity float64 = 0.00
	var europeanVelocity float64 = 0
	var americanVelocity float64 = 0
	fmt.Scan(&valueVelocity)
	fmt.Printf("Вы ввели %.2f м/с\n", valueVelocity)
	europeanVelocity = valueVelocity * 3.6
	fmt.Printf("Тело движется со скоростью %.2f км/ч\n", europeanVelocity)
	americanVelocity = valueVelocity * 2.237
	fmt.Printf("Тело движется со скоростью %.2f км/ч\n", americanVelocity)
}
