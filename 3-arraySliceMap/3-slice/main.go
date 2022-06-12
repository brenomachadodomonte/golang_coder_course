package main

import (
	"fmt"
	"reflect"
)

func main(){

	a1 := [3]int{1, 2, 3} // an array (fixed size)
	s1 := []int{1, 2, 3} // a slice	(dinamic size)

	fmt.Println(a1, s1);
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1));

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice is not an array, Slice defines part of an array
	s2 := a2[1:3]
	fmt.Println(a2, s2);

	s3 := a2[:2]
	fmt.Println(a2, s3);

	s4 := s2[:1]
	fmt.Println(s2, s4);

	// will change the original array
	s4[0] = 15
	fmt.Println(a2);
}