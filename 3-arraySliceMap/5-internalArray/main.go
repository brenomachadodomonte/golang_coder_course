package main

import "fmt"

func main(){

	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	
	fmt.Println(s1, s2) 

	s1[0] = 7 // will change slice 1 and slice 2 at index 0
	fmt.Println(s1, s2) 

}