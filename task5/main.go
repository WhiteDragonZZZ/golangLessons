package main

import "fmt"

func main() {
	dayOfWeek := make([]string, 7, 7)
	dayOfWeek = []string{
		"Monday",
		"Thuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	workDay := dayOfWeek[:5]
	fmt.Printf("Need slice %q,а также длина %d и объем %d \n", workDay, len(workDay), cap(workDay))
	relaxDay := dayOfWeek[5:7]
	fmt.Println("Need slice %q,а также длина %d и объем %d", relaxDay, len(relaxDay), cap(relaxDay))
}
