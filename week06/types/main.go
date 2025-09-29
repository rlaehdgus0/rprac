package main

import (
	"fmt"
	"reflect"
)

func main() {
	//대문자는 외부에서 참조 가능
	var f64 float64
	totalPrice := 1000

	fmt.Println(totalPrice)
	fmt.Println(f64, reflect.TypeOf(f64))
}
