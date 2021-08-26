package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var total int
	for _, val := range s.score{
		total += val
	}
	return float64(total/len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for idx, val := range s.score{
		if min > val{
			min = val
			name = s.name[idx]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for idx, val := range s.score{
		if max < val{
			max = val
			name = s.name[idx]
		}
	}
	return max, name
}



func main() {
	var a = Student{}
	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input " + strconv.Itoa(i+1) + " Student’s Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
