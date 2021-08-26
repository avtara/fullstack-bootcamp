package main

import "fmt"

func main() {
	// input
	var studentScore int = 80
	var studentName string

	fmt.Scanf("%s", &studentName)
	fmt.Scanf("%d", &studentScore)

	if studentScore > 80 && studentScore <= 100 {
		fmt.Println("Grade", studentName, ":", "A")
	} else if studentScore >= 65 && studentScore < 80 {
		fmt.Println("Grade", studentName, ":", "B")
	} else if studentScore >= 50 && studentScore < 65{
		fmt.Println("Grade", studentName, ":", "C")
	} else if studentScore >= 0 && studentScore < 35{
		fmt.Println("Grade", studentName, ":", "D")
	} else {
		fmt.Println("Nilai", studentName, ":", "Invalid")
	}
}