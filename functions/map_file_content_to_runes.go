package Ascii

import (
	"log"
	"strings"
)

// this function will split the content content file and build map to each character
func MapFileContentToRunes( fileContent string) map[rune][]string {
	startChar := ' '
	// blocks := strings.Split(fileContent, "\r\n\r\n")
	// if(fileName)
	blocks := strings.Split(fileContent, "\n\n")

	// for _, v := range fileContent {
	// 	fmt.Printf("v: %v\n", v)
	// }
	charMap := make(map[rune][]string)
	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		if len(lines) > 0 {
			char := rune(startChar + rune(i))
			
			charMap[char] = lines
		} else {
			log.Fatal("Worning: empty or malformed block at index ", i)
		}
	}

	return charMap
}
