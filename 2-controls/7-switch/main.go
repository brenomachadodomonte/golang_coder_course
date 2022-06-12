package main

import (
	"fmt"
	"time"
)

func returnType(i interface{}) string {
	
	switch i.(type) {
	case int: 
		return "Integer"
	case float64, float32: 
		return "Decimal Number"
	case string: 
		return "String"	
	case func():
		return "Function"
	default:
		return "Type unknown"
	}
}

func main(){

	fmt.Println(returnType(2.3));
	fmt.Println(returnType(4));
	fmt.Println(returnType("Hello"));

	fmt.Println(returnType(func() {}));
	fmt.Println(returnType(time.Now()));

}