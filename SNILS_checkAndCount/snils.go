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
		err := errors.New("неверная длинна СНИЛС")
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

	// 112233445 - snils
	// 012345678 - i
	// 987654321 - multiplier
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
