package snils

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func CheckAndCount(snils string) (bool, error) {

	// ADD: Проверка контрольного числа Страхового номера проводится только для номеров больше номера 001-001-998.

	if len(snils) != 14 {
		err := errors.New("неверная длина СНИЛС")
		return false, err
	}

	/*
	   Есть ли готовая возможность задать маску для строки и проверять по ней
	   по аналогии с шаблоном для формата даты в пакете time?
	*/
	for _, elem := range snils {
		if unicode.IsLetter(elem) {
			err := errors.New("недопустимый символ в СНИЛС")
			return false, err
		}
	}

	return count(snils), nil
}

func count(snils string) bool {

	snilsLeft := snils[:11]
	snilsLeft = strings.ReplaceAll(snilsLeft, "-", "")

	var sum int

	controlSum, _ := strconv.Atoi(string(snils[12:]))

	// fmt.Println("snilsLeft", snilsLeft) // 112233445
	//									 i    012345678
	// 							multiplier    987654321
	for i, value := range snilsLeft {

		sum += int(value-'0') * (9 - i)
	}

	sumCheck := sum % 101

	switch sumCheck {
	case 100:
		sumCheck = 0
		fallthrough
	case controlSum:
		return true
	}

	return false
}
