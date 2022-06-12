package main

import "fmt"

func main(){

	var a int		// default 0
	var b float64	// default 0
	var c bool		// default false
	var d string	// default ""
	var e *int		// default <nil>

	fmt.Printf("%v %v %v %v", a, b, c, d, e)

}