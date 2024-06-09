package Ascii

import (
	"fmt"
	"os"
)

// this function will read the content of the file and retuns it as a string
func ReadShapeFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error in opening the file ", err)
	}
	// skipping the first byte which is empty and retuns a string starting from the second
	return string(file[1:])
}
