package main

import "fmt"

// there are no ternary operator in go :(
func getResult(grade float64) string {
	if(grade >= 7){
		return "Approved"
	}

	return "Failed";
}

func main() {
	
	fmt.Println(getResult(7.2))

}

