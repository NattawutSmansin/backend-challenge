package main

import (
	firstChallenge "backend-challenge/module/firstChallenge"
	"backend-challenge/router"
	secondChallenge "backend-challenge/module/secondChallenge"
	"fmt"
)

func main() {
	fmt.Println("----------------- 1. จงหาเส้นทางที่มีค่ามากที่สุด -------------------")
	firstChallenge.FirstChallenge()

	fmt.Println("----------------- 2. จับฉันให้ได้สิ ซ้าย-ขวา-เท่ากับ -------------------")
	inputStrings := "LLRR="
	secondChallenge.SecondChallenge(inputStrings)

	fmt.Println("----------------- 3. พาย ไฟ ได - Pie Fire Dire -------------------")
	router.StartGRPCServer()
}
