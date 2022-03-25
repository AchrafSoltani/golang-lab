package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func sayHello(s string) string {
	return "Hello " + s
}

func addition(x int, y int) int {
	return x + y
}

func createArray() [4]string {
	var beatles [4]string
	beatles[0] = "John"
	beatles[1] = "Paul"
	beatles[2] = "Ringo"
	beatles[3] = "George"

	return beatles
}

func main() {
	fmt.Println(sayHello("Ash"))
	fmt.Print("1 + 2 = ")
	fmt.Println(addition(1, 2))
	fmt.Print("And that is ")
	var truth bool = true
	fmt.Println(truth)
	var data = createArray()
	fmt.Print("Variable data is of type ")
	fmt.Println(reflect.TypeOf(data))
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))

}
