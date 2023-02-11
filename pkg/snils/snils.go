/*
Package snils Realising Algorithm of checking the control number of the Insurance number (SNILS).
According to http://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/.
Manual check SNILS http://portal.fss.ru:8585/fss/lgot/snils.
Play code https://go.dev/play/p/s6cCUxK7ekU.
*/
package snils

import (
	"fmt"
	"regexp"
	"strconv"
)

// Checker interface implements CheckAndCount() error.
type Checker interface {
	CheckAndCount() error
}

// type Snils contains string value
type Snils string

/*
Method CheckAndCount checks regexp validity and SNILS validity to regexp. Returns wrapped error.
Otherwise (err == nil) initiate method count() which counts SNILS and compares it to checkSum returning wrapped error or nil.
*/
func (s Snils) CheckAndCount() error {

	regexpStr := "^(\\d{3})[ -]*(\\d{3})[ -]*(\\d{3})[ -]*(\\d{2})\\s*$"

	re, err := regexp.Compile(regexpStr) // checking regexp validity
	if err != nil {
		return fmt.Errorf("failed compile regex: %s", regexpStr)
	}

	matched := re.MatchString(string(s)) // checking snils validity
	if !matched {
		return fmt.Errorf("invalid SNILS format: %s", s)
	}

	re.FindStringSubmatch(string(s)) // getting string slice of regexp groups (len == 5)

	checkSum := re.FindStringSubmatch(string(s))[4]
	snils := re.FindStringSubmatch(string(s))[1] + re.FindStringSubmatch(string(s))[2] + re.FindStringSubmatch(string(s))[3]

	if snilsInt, _ := strconv.Atoi(snils); snilsInt <= 1001998 {
		return fmt.Errorf("checking SNILS is carried out only for numbers larger than the number 001-001-998")
	}

	return s.count(snils, checkSum)
}

// Method count counts SNILS and compares it to checkSum. Returns wrapped error or nil.
func (s Snils) count(snils, checkSum string) error {
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
		return nil
	}

	return fmt.Errorf("SNILS doesn't match. SNILS: %s", s)
}

// Func checkSnils can be used to mock interface
func checkSnils(snils Checker) error {
	return snils.CheckAndCount()
}
