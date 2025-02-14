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
	var skip int = 999

	for i, v := range sliceStr {

		if i == 0 && isDigit(v) {
			return "", ErrInvalidString
		}

		if i+1 < len(sliceStr) && isDigit(sliceStr[i+1]) {
			multi := int(sliceStr[i+1] - '0')
			skip = i + 1

			for j := 0; j < multi; j++ {
				result += string(v)
			}
		} else if i != skip {
			result += string(v)
		}
	}

	return result, nil
}
