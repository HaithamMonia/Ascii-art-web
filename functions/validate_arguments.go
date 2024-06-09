package Ascii

import (
	"os"
)

func ValidateArugments() bool {
	//file.txt
	fileName := os.Args[2]
	txtCheck := len(os.Args[3]) - 4

	if txtCheck <= 0 {
		return false
	}
	if fileName == "standard.txt" || fileName == "shadow.txt" || fileName == "thinkertoy.txt" {
		if os.Args[3][txtCheck:] == ".txt" {
			return true
		}
	}
	return false
}
