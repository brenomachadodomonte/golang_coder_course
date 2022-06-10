package main

import "fmt"

func getGradeResult(grade float64) string {
	var g = int(grade)
	switch g {
	case 10: 
		fallthrough // go to the next case
	case 9: 
		return "A"
	case 8, 7: 
		return "B"	
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid grade"
	}
}

func getGradeResult2(grade float64) string {
	switch {
	case grade >= 9: 
		return "A"
	case grade >= 8:
		return "B"
	case grade >= 5:
		return "C"
	case grade >= 3:
		return "D"
	default:
		return "E"		
	}
}

func main(){
	fmt.Println(getGradeResult(9.8));
	fmt.Println(getGradeResult2(7.1));
	fmt.Println(getGradeResult2(3.5));
}