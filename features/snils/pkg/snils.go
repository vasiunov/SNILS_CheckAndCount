package snils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
	Realising Algorithm for generating the control number of the Insurance number
	According to http://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/
	Manual cheking SNILS http://portal.fss.ru:8585/fss/lgot/snils
*/

func main() {

	var Snils = "112-233-445 95" // checking string

	fmt.Println(CheckAndCount(Snils)) // result
}

func CheckAndCount(snils string) (bool, error) {

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

	var sum, controlSum int

	controlSum, _ = strconv.Atoi(string(snils[12:]))

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
