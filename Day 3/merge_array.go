package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	rawArr := append(arrayA, arrayB...)
	unique := make(map[string]bool, len(rawArr))
	resArr := make([]string, len(unique))
	for _, elem := range rawArr {
		if len(elem) != 0 {
			if !unique[elem] {
				resArr = append(resArr, elem)
				unique[elem] = true
			}
		}
	}
	return resArr
}



func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// [“king”, “devil jin”, “akuma”, “eddie”, “steve”, “geese”]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// [“sergei”, “jin”, “steve”, “bryan”]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// [“alisa”, “yoshimitsu”, “devil jin”, “law”]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// [“devil jin”, “sergei”]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// [“hwoarang”]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
