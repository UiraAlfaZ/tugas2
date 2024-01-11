package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//string ke float
	name := "Wira Abdi Fazriansyah"
	num, err := strconv.ParseFloat(name, 64)
	if err != nil {
		fmt.Println("Gagal mengonversi string ke float")
	} else {
		fmt.Println(num, reflect.TypeOf(name))
	}
	//int ke string
	age := 22
	str := strconv.Itoa(age)
	fmt.Println(age, reflect.TypeOf(age))

	//float ke string
	height := 123.45
	str = strconv.FormatFloat(height, 'f', 2, 64) //dengan = bukan :=
	fmt.Println(str, reflect.TypeOf(str))
}
