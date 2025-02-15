package hw02unpackstring

import (
	"errors"
	"regexp"
)

var ErrInvalidString = errors.New("invalid string")

func isDigit(r rune) bool {
	re := regexp.MustCompile(`\d`)
	return re.MatchString(string(r))
}

func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	sliceStr := []rune(input)
	var result string

	for i := 0; i < len(sliceStr); {
		if i == 0 && isDigit(sliceStr[i]) {
			return "", ErrInvalidString
		}
		if isDigit(sliceStr[i]) && (i+1) < len(sliceStr) {
			if isDigit(sliceStr[i+1]) {
				return "", ErrInvalidString
			}
		}
		if i+1 < len(sliceStr) && isDigit(sliceStr[i+1]) {
			multi := int(sliceStr[i+1] - '0')

			for j := 0; j < multi; j++ {
				result += string(sliceStr[i])
			}
			i++
		} else {
			result += string(sliceStr[i])
		}
	}

	return result, nil
}
