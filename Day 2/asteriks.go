package main

import "fmt"

func playWithAsterik(n int) {
	var space = n - 1
	for i := 0; i < n; i++{
		for j := 0 ; j < space; j++{
			fmt.Print(" ")
		}
		space -= 1
		for j := 0 ; j <= i; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}



func main() {

	playWithAsterik(5)

	/*

	       *

	      * *

	     * * *

	    * * * *

	   * * * * *

	*/

}