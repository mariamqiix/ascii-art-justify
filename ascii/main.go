package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

// colorWJL
func main() {
	validation := ascii.Validation()

	fmt.Println(validation)
	fmt.Println(len(os.Args))
	index := 1

	if !strings.Contains(validation, "W") && validation != "yes" {
		index++
	} else if !strings.Contains(validation, "F") {
		index = len(os.Args) - 1
	} else if strings.Contains(validation, "F") {
		index = len(os.Args) - 2
	}

	ascii.CheckLetter(os.Args[index])
	// count := strings.Count(os.Args[index], " ")
	WordsInArr := strings.Split(os.Args[index], "\\n")
	fileName := "standard"
	if (len(os.Args) == 3 && validation == "yes") || (len(os.Args) == 4 && !strings.Contains(validation, "W")) || (strings.Contains(validation, "F")) {
		fileName = strings.ToLower(os.Args[index+1])
	}
	
	if !ascii.CheckFont(fileName) {
		ascii.Error()
	}

	if !(ascii.CheckTextSizeWithWidth(WordsInArr,fileName)) {
		fmt.Println("too much words , write less")
		return
	}
	if ascii.OnlyContains(os.Args[index], "\\n") {
		WordsInArr = WordsInArr[:len(WordsInArr)-1]
	}

	FirstWord := true

	for l := 0; l < len(WordsInArr); l++ {
		var Words [][]string
		Text1 := strings.ReplaceAll(WordsInArr[l], "\\t", "   ")
		for j := 0; j < len(Text1); j++ {
			Words = append(Words, ascii.ReadLetter(Text1[j], fileName))
		}
		if len(Text1) != 0 {
			count := strings.Count(Text1," ")
			align := "left"
			if strings.Contains(validation, "J") {
				align = ascii.ReturnAlign()
			}
			if strings.Contains(validation, "output") {
				if strings.Contains(validation, "C") && l == 0 {
					fmt.Println("The coloerd text can't be print inside the file")
				}
				ascii.WriteFile(Words, FirstWord, validation ,align, count)
				FirstWord = false
			} else if strings.Contains(validation, "color") {
				letter1 := "NO!!-"
				color := ascii.CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[1], "--color=")))
				if strings.Contains(validation, "22") {
					color = ascii.CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[2], "--color=")))
				}
				if validation == "colorW2L" || validation == "colorW2LF" || validation == "colorWLJ" || validation == "colorWLJF" {
					letter1 = os.Args[2]
				} else if strings.Contains(validation, "L") {
					letter1 = os.Args[index-1]
				}
				if strings.Contains(validation, "J") {
					if validation == "outputWJ" {
							align = strings.ToLower(strings.TrimPrefix(os.Args[2], "--align="))
					} else {
							align = strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))

					}
				}
				ascii.PrintWithColor(Words, color, Text1, letter1, validation , "center", count)
			} else {
				for w := 0; w < 8; w++ {
					for n := 0; n < len(Words); n++ {
						if validation == "justify" {
							align = strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))
							fmt.Print(ascii.PrintWithJustify(Words, align, n, w,count))
						} else {
							fmt.Print(Words[n][w])
						}
					}
					if w+1 != 8 {
						fmt.Println()
					}
				}
				fmt.Println()
			}
		}
	}
}
