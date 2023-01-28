package snils

import (
	"errors"
	"regexp"
	"strconv"
)

type Checker interface {
	CheckAndCount() (bool, error)
}

type Snils string

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

func checkSnils(snils Checker) (bool, error) {
	return snils.CheckAndCount()
}
