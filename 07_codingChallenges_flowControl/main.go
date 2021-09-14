package main

import "fmt"

func main() {

	//ex6
	age := 18
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}

	// //ex5
	// i := 1995
	// for i <= 2021 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// //ex4
	// for i := 1; i < 500; i++ {
	// 	if i%7 == 0 && i%5 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// //ex3
	// j := 0
	// for i := 1; i < 50; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Println(i)
	// 		j++
	// 	}
	// 	if j == 3 {
	// 		break
	// 	}
	// }

	// //ex2
	// for i := 1; i < 50; i++ {
	// 	if i%7 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// //ex1
	// for i := 1; i < 50; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
}
