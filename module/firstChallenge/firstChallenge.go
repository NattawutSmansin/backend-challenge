package firstChallenge

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var memo map[[2]int]int // ใช้เก็บค่าที่คำนวณแล้ว
func FirstChallenge() {
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching the file: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading the response body: ", err)
	}

	var dataJsons [][]int
	err = json.Unmarshal(body, &dataJsons)
	if err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	memo = make(map[[2]int]int)

	var row, col int = 0, 0
	result := maxPath(dataJsons, row, col)
	fmt.Println("ผลรวมค่ามากที่สุด:", result)

	return

}

func maxPath(dataJsons [][]int, row, col int) int {
	if row == len(dataJsons)-1 {
		return dataJsons[row][col]
	}

	key := [2]int{row, col}
	if val, found := memo[key]; found {
		return val
	}

	left := maxPath(dataJsons, row+1, col)
	right := maxPath(dataJsons, row+1, col+1)

	best := dataJsons[row][col] + max(left, right)

	memo[key] = best
	return best
}
