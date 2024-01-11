package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Wira Abdi Fazriansyah"
	age := 22
	heigth := 1.63
	isMarried := false
	interestInCoding := true

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(age, reflect.TypeOf(age))
	fmt.Println(heigth, reflect.TypeOf(heigth))
	fmt.Println(isMarried, reflect.TypeOf(isMarried))
	fmt.Println(interestInCoding, reflect.TypeOf(interestInCoding))
}
