package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	var Snils = "112-233-445 95" // checking string

	fmt.Println(checkAndCountSNILS(Snils)) // result
}

func checkAndCountSNILS(snils string) (bool, error) {

	if len(snils) != 14 {
		err := errors.New("неверная длина СНИЛС")
		return false, err
	}

	/*
	   Есть ли готовая возможность задать маску для строки и проверять по ней?
	   По аналогии с шаблоном для даты в пакете time
	*/
	for _, elem := range snils {
		if unicode.IsLetter(elem) {
			err := errors.New("недопустимый символ в СНИЛС")
			return false, err
		}
	}

	return countSnils(snils), nil
}

func countSnils(snils string) bool {

	snilsLeft := snils[:11]
	snilsLeft = strings.ReplaceAll(snilsLeft, "-", "")

	// fmt.Println("snilsLeft", snilsLeft) // 112233445
	//									 i    012345678
	// 							multiplier    987654321

	var sum, controlSum int

	controlSum, _ = strconv.Atoi(string(snils[12:]))

	for i, value := range snilsLeft {

		sum += int(value-'0') * (9 - i)
	}

	if sum == controlSum {
		return true
	}

	return false
}
