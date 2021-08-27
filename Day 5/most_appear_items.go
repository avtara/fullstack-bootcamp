package main

import (
	"fmt"
	"sort"
)


func MostAppearItem(items []string) string {
	var res string
	var count = make(map[string]int)
	var keyCollection []int
	var switchMapCount = make(map[int][]string)

	// save to map ex: map[golang:1 js:4 ruby:2]
	for _,val := range items{
		count[val]++
	}

	// swap key and value in map ex: map[1:[golang] 2:[ruby] 4:[js]]
	for idx, val := range count {
		switchMapCount[val] = append(switchMapCount[val], idx)
	}

	// save total items to slice
	for idx := range switchMapCount {
		keyCollection = append(keyCollection, idx)
	}

	//  sort slice
	sort.Ints(keyCollection)

	// print / add item to res using switchMapCount by sorted key / keyCollection
	for _, val := range keyCollection {
		for _, value := range switchMapCount[val] {
			res += fmt.Sprintf("%s->%d ", value, val)
		}
	}

	return res
}



func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1 basketball->1 tenis->1
}