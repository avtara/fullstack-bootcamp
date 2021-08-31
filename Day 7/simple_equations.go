package main

import "fmt"

func SimpleEquations(a, b, c int) {
	var tempA, tempB, tempC = 1, 1, 1
	for ; a - (tempA + tempB + tempC ) >= 0 || b - (tempA * tempB * tempC ) >= 0 || c - (tempA * tempA+ tempB*tempB + tempC*tempC ) >= 0; tempA++{
		for ; a - (tempA + tempB + tempC ) >= 0 || b - (tempA * tempB * tempC ) >= 0 || c - (tempA * tempA+ tempB*tempB + tempC*tempC ) >= 0; tempB++{
			for ; a - (tempA + tempB + tempC ) >= 0 || b - (tempA * tempB * tempC ) >= 0 || c - (tempA * tempA+ tempB*tempB + tempC*tempC ) >= 0; tempC++{
				if a - (tempA + tempB + tempC ) == 0 && b - (tempA * tempB * tempC ) == 0 && c - (tempA * tempA+ tempB*tempB + tempC*tempC ) == 0{
					fmt.Println(tempA,tempB,tempC)
					return
				}
			}
			tempC = 1
		}
		tempB = 1
	}
	fmt.Println("no solution")
}



func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
	SimpleEquations(6, 6, 14) //1 2 3
	SimpleEquations(3, 1, 3) //1 1 1
	SimpleEquations(4, 2, 6) //1 1 2
	SimpleEquations(5, 3, 11) //1 1 3
	SimpleEquations(5, 4, 9) //1 2 2
	SimpleEquations(6, 4, 18) //1 1 4
	SimpleEquations(6, 6, 14) //1 2 3
	SimpleEquations(6, 8, 12) //2 2 2
	SimpleEquations(7, 5, 27) //1 1 5
	SimpleEquations(7, 8, 21) //1 2 4
	SimpleEquations(7, 9, 19) //1 3 3
	SimpleEquations(7, 12, 17) //2 2 3
	SimpleEquations(8, 6, 38) //1 1 6
	SimpleEquations(8, 10, 30) //1 2 5
	SimpleEquations(8, 12, 26) //1 3 4
	SimpleEquations(8, 16, 24) //2 2 4
	SimpleEquations(8, 18, 22) //2 3 3
	SimpleEquations(9, 12, 41) //1 2 6
	SimpleEquations(9, 15, 35) //1 3 5
	SimpleEquations(9, 16, 33) //1 4 4
	SimpleEquations(9, 20, 33) //2 2 5
	SimpleEquations(9, 24, 29) //2 3 4
	SimpleEquations(9, 27, 27) //3 3 3
	SimpleEquations(10, 18, 46) //1 3 6
	SimpleEquations(10, 20, 42) //1 4 5
	SimpleEquations(10, 24, 44) //2 2 6
	SimpleEquations(10, 30, 38) //2 3 5
	SimpleEquations(10, 32, 36) //2 4 4
	SimpleEquations(10, 36, 34) //3 3 4
	SimpleEquations(11, 36, 49) //2 3 6
	SimpleEquations(11, 40, 45) //2 4 5
	SimpleEquations(11, 45, 43) //3 3 5
	SimpleEquations(11, 48, 41) //3 4 4
}