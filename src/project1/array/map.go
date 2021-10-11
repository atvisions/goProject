package main

import (
	"errors"
	"fmt"
	"os"
)

type s1 []int
type s2 []s1

var searchErr error = errors.New("search error")

func main() {
	a := s2{
		s1{1, 2, 3, 4, 5, 6},
		s1{7, 8, 9, 10, 11, 12},
		s1{13, 14, 15, 16, 17, 18},
		s1{20, 21, 22, 23, 24, 25}}

	result, err := search(a[:][:], 4, 6, 21)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if result == true {
		fmt.Println(a)
	}
}

func search(a s2, row, col, key int) (flag bool, err error) {
	if row < 0 || col < 0 {
		return false, searchErr
	}
	flag = false
	rowNum := 0
	colNum := col - 1
	for rowNum < row && colNum >= 0 {
		if a[rowNum][colNum] == key {
			flag = true
			break
		} else if a[rowNum][colNum] > key {
			colNum--
		} else {
			rowNum++
		}
	}
	return
}