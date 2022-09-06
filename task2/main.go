package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
Задание 2
Ввести первое строковое числовое значение '154'
Ввести второе число '90'
Вывести сумму строки и числа '244'
*/

func main() {

	var stringValue string
	fmt.Scan(&stringValue)

	var intValue int = 0
	fmt.Scan(&intValue)
	result, err := strconv.Atoi(stringValue)
	if err != nil {
		log.Fatal(err)
	}
	result = result + intValue
	fmt.Println(result)
}
