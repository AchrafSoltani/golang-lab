package main

import "fmt"

const lol string = "lots of laughs"

func pointer() {
	s := "pointer string"
	fmt.Print("Pointer address : ")
	fmt.Println(&s)
	showMemAddress(&s)
}

func main() {
	var s, t string = "hi", "thanks"
	fmt.Println(s)
	fmt.Println(t)
	r := "short variable form" // short variable form
	fmt.Println(r)
	pointer()
}

func showMemAddress(x *string) {
	fmt.Println(x)
	fmt.Print("getting the value from a pointer : ")
	fmt.Println(*x)
	fmt.Println(lol)
}
