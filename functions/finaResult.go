package Ascii

import (
	"strings"
)

func FinalResult(text string, font string) string {
	//extracting the file name from the args
	var shapeFileContent string
	// check what's the defualt ?
	// if len(os.Args)==4 {
	// 	if len(os.Args) == 2 { // standard is the default
	// 		shapeFileContent = functions.ReadShapeFile("standard.txt")
	// 	} else

	shapeFileContent = ReadShapeFile(font)

	charactersMap := MapFileContentToRunes(shapeFileContent)
	textToDraw := text
	// This will help in handling \n each index in textToDrawSplit will be considered a line
	textToDrawSplit := strings.Split(textToDraw, `\n`)
	result := ""
	for i, line := range textToDrawSplit {
		if line != "" {
			generateShape := GenerateShape(line, charactersMap)
			//fmt.Println(generateShape)
			result += generateShape

		}

		if i < len(textToDrawSplit)-1 {
			result += "\n"
		}
	}
	return result
}
