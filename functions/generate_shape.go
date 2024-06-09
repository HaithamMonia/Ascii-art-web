package Ascii

import (
	"strings"
)

// generates the art and returns it as a string
func GenerateShape(text string, charMap map[rune][]string) string {
	var result []string
	//initialization
	for i := 0; i < 8; i++ {
		result = append(result, "")
	}

	for _, char := range text {
		charShape, found := charMap[char]
		if found {
			for i, line := range charShape {

				result[i] += line
			}
		}
	}
	return strings.Join(result[:len(result)-1], "\n")
}
