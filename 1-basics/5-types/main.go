package main

import (
	"fmt"
	"reflect"
	"math"
)

func main(){

	// integer numbers
	fmt.Println(1, 2, 1000)
	fmt.Println("The number is", reflect.TypeOf(32000))

	// unsigned uint8 uint16 uint32 uint64
	var b byte = 255 //uint8
	fmt.Println("The b is", reflect.TypeOf(b))

	// signed
	i1 := math.MaxInt64
	fmt.Println("Int max value is", i1)
	fmt.Println("The b is", reflect.TypeOf(i1))

	var i2 rune = 'a' // unicode (int32)
	fmt.Println("The rune is", reflect.TypeOf(i2))
	fmt.Println("The i2 value is", i2)

	// decimal numbers
	var x float32 = 49.99 // or x := float32(49.99) 
	fmt.Println("The x type is", reflect.TypeOf(x))
	fmt.Println("The literal type is", reflect.TypeOf(49.99))


	// boolean 
	bo := true
	fmt.Println("The bo type is", reflect.TypeOf(bo))

	// string
	s1 := "Hi, my name is Breno"
	fmt.Println(s1 + "1")
	fmt.Println("The string length is", len(s1))

	// multiline string
	s2 := `Hello
	 this
	 is
	 a 
	 multiline string
	`
	fmt.Println(s2)
	fmt.Println("The string length is", len(s2))

	// char???
	char := 'a' // is int32
	fmt.Println("The char type is", reflect.TypeOf(char))


}