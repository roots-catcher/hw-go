package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	sliceStr := []rune(input)
	var result string

	for i := 0; i < len(sliceStr); i++ {
		if i == 0 && unicode.IsDigit(sliceStr[i]) {
			return "", ErrInvalidString
		}
		if i+1 < len(sliceStr) && unicode.IsDigit(sliceStr[i+1]) {
			multi := int(sliceStr[i+1] - '0')

			for j := 0; j < multi; j++ {
				result += string(sliceStr[i])
			}
			i++
		} else {
			result += string(sliceStr[i])
		}
		if unicode.IsDigit(sliceStr[i]) && (i+1) < len(sliceStr) {
			if unicode.IsDigit(sliceStr[i+1]) {
				return "", ErrInvalidString
			}
		}
	}

	return result, nil
}
