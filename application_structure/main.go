package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello Go World!!")
	distance := 60.8

	fmt.Printf("The distance in miles is %f \n", distance*0.62137)

	text := 5
	fmt.Printf("Sum: %f", 2.7+float32(text))
}
