package main

import "fmt"

func getPrize() (int, string) {
	i := 3
	s := "Supercar"
	return i, s
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	quantity, prize := getPrize()
	fmt.Printf("You won %v %v\n", quantity, prize)
	result := sumNumbers(1, 2, 3, 4, 5)
	fmt.Println(result)
	fmt.Println(namedReturn())
	recursive(1, 0)
	fn := func() {
		fmt.Println("function called")
	}
	fn()
}

func namedReturn() (x, y string) {
	x = "hi "
	y = " yo"
	return
}

func recursive(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full! I've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm hungry! I've eaten %d\n", eaten)
	return recursive(portion, eaten)
}
