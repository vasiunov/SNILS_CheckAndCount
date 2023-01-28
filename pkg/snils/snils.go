/*
Package snils Realising Algorithm of checking the control number of the Insurance number (SNILS).
According to http://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/.
Manual check SNILS http://portal.fss.ru:8585/fss/lgot/snils.
Play code https://go.dev/play/p/IMrxtpMOagJ.
*/
package snils

import (
	"errors"
	"regexp"
	"strconv"
)

// Checker interface implements CheckAndCount() (bool, error)
type Checker interface {
	CheckAndCount() (bool, error)
}

// type Snils contains string value
type Snils string

/*
CheckAndCount checks regexp validity and SNILS validity to regexp.
Returns false and error (containing result of checking).
Otherwise initiate method count() which counts SNILS and compares it to checkSum returning bool
*/
func (s Snils) CheckAndCount() (bool, error) {

	re, err := regexp.Compile(`^(\d{3})[ -]*(\d{3})[ -]*(\d{3})[ -]*(\d{2})\s*$`) // checking regexp validity
	if err != nil {
		return false, err
	}

	matched := re.MatchString(string(s)) // checking snils validity
	if !matched {
		return matched, errors.New("invalid SNILS format") // CODE REVIEW NEEDED
	}

	re.FindStringSubmatch(string(s)) // getting string slice of regexp groups (len == 5)

	checkSum := re.FindStringSubmatch(string(s))[4]
	snils := re.FindStringSubmatch(string(s))[1] + re.FindStringSubmatch(string(s))[2] + re.FindStringSubmatch(string(s))[3]

	return s.count(snils, checkSum), nil
}

// Method count counts SNILS and compares it to checkSum
func (s Snils) count(snils, checkSum string) bool {
	var sum int

	checkSumInt, _ := strconv.Atoi(checkSum)

	// 112233445 - snils
	// 012345678 - i
	// 987654321 - multiplier
	for i, value := range snils {
		sum += int(value-'0') * (9 - i)
	}

	switch sum % 101 {
	case 100:
		sum = 0
		fallthrough
	case checkSumInt:
		return true
	}

	return false
}

// Func checkSnils can be used to mock interface
func checkSnils(snils Checker) (bool, error) {
	return snils.CheckAndCount()
}
