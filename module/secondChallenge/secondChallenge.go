package secondchallenge

import (
	"fmt"
	"strconv"
	"strings"
)

var result string

func SecondChallenge(input string) {
	patterns := strings.Split(input, "")
	if len(patterns) < 2 || len(patterns) == 0 {
		fmt.Println("จงใส่ข้อมูล Pattern ดังนี้ : `= หรือ L หรือ R` มากกว่าหรือเท่ากับ 2")
	}

	numberSlices := []int{}
	index := 0

	ascendOrder(index, numberSlices, patterns)
	fmt.Println("ผลรวมค่าน้อยที่สุด:", result)
}

func ascendOrder(index int, numberSlices []int, patterns []string) {
	countPattern := len(patterns) + 1

	if len(numberSlices) == countPattern {
		tempResult := convertToString(numberSlices)
		if result == "" || tempResult < result { // เก็บค่าเฉพาะผลลัพธ์ที่น้อยที่สุด
			result = tempResult
		}
		return
	}

	for num := 0; num <= 9; num++ {
		newSlice := append([]int{}, numberSlices...) // copy numberSlices
		newSlice = append(newSlice, num)

		if index == 0 {
			ascendOrder(index+1, newSlice, patterns)
		} else {
			if getConditionBySymbol(index, num, patterns, numberSlices) {
				ascendOrder(index+1, newSlice, patterns)
			}
		}
	}
}


func getConditionBySymbol(index, num int, patterns []string, numberSlices []int) bool {
	if len(numberSlices) < 1 {
		return false
	}

	lastNum := numberSlices[len(numberSlices)-1]
	switch patterns[index-1] {
	case "L":
		return lastNum > num
	case "R":
		return lastNum < num
	case "=":
		return lastNum == num
	}
	return false
}

func convertToString(numbers []int) string {
	str := ""
	for _, num := range numbers {
		str += strconv.Itoa(num)
	}
	return str
}
