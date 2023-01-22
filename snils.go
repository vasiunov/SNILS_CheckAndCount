package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var Snils = "112-233-445 95"

	checkSNILS(Snils)

	fmt.Println(checkSNILS(Snils))
}

func checkSNILS(snils string) bool {
	snilsLeft := snils[:11]
	snilsLeft = strings.ReplaceAll(snilsLeft, "-", "")

	fmt.Println("snilsLeft", snilsLeft) // 112233445
	var sum, controlSum, intLeft, intLeftResult, intRight, intRightResult int

	for i := 0; i <= 4; i++ {
		if i == 4 {
			intEl, _ := strconv.Atoi(string(snils[9-i+1]))
			sum += intEl * (9 - i)

			controlSum, _ = strconv.Atoi(string(snils[12:]))

			// fmt.Println("sum =", sum)
			// fmt.Println("controlSum =", controlSum)

			if sum == controlSum {
				return true
			}

		} else {

			intLeft, _ = strconv.Atoi(string(snilsLeft[i]))
			intLeftResult = intLeft * (9 - i)

			intRight, _ = strconv.Atoi(string(snilsLeft[len(snilsLeft)-1-i]))
			intRightResult = intRight * (1 + i)

			sum += intLeftResult + intRightResult

			// fmt.Println("sum =", sum)
			// fmt.Println("i =", i)
		}
	}
	return false
}
