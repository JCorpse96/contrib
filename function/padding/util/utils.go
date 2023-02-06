package util

import (
	"strings"
)

func PaddingRight(value string, length int, padCharacter string) string {
	var padCountInt = 1 + ((length - len(padCharacter)) / len(padCharacter))
	var retStr = value + strings.Repeat(padCharacter, padCountInt)

	return retStr[:length]
}
