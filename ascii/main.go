package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	validation := ascii.Validation()
	fmt.Println(validation)
	index := 1
	if validation == "output" || validation == "color" || validation == "justify" {
		index++
	} else if validation == "colorWletter" || validation == "colorWletterWfont" || validation == "outputWC" || validation == "outputWCF" {
		index += 2
	} else if validation == "colorW2letter" {
		index += 4
	} else if validation == "outputWCL" || validation == "outputWCLF" {
		index += 3
	}

	WordsInArr := strings.Split(os.Args[index], "\\n")

	fileName := "standard"
	if len(os.Args) == 3 && validation != "output" && validation != "color" && validation != "justify" {
		fileName = strings.ToLower(os.Args[index+1])
	} else if len(os.Args) == 4 && validation != "colorWletter" && validation != "outputWC" {
		fileName = strings.ToLower(os.Args[index+1])
	} else if validation == "colorWletterWfont" || validation == "outputWCF" {
		fileName = strings.ToLower(os.Args[4])
	} else if validation == "outputWCLF" {
		fileName = strings.ToLower(os.Args[5])
	}

	if ascii.CheckPrint(WordsInArr, -1, fileName) == -1 {
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
		if strings.Contains(validation, "output") {
			if strings.Contains(validation, "C") {
				fmt.Println("We can't print color in the file")
			}
			ascii.WriteFile(Words, FirstWord)
			FirstWord = false
		} else if validation == "color" || validation == "colorWletter" || validation == "colorWletterWfont" || validation == "colorW2letter" {
			letter1 := "NO!!-"
			color := ascii.CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[1], "--color=")))
			if validation == "colorWletter" || validation == "colorWletterWfont" {
				letter1 = os.Args[index-1]
			} else if validation == "colorW2letter" {
				letter1 = os.Args[2]
			}
			ascii.PrintWithColor(Words, color, Text1, letter1, validation)
			// else if validation == "justify" {
			// 	align := strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))
			// 	ascii.PrintWithJustify(Words, WordsInArr, align, Text1, fileName, l)
			// }
		} else {
			for w := 0; w < 8; w++ {
				if len(Text1) == 0 {
					break
				}
				for n := 0; n < len(Words); n++ {
					if validation == "justify" {
						align := strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))
						ascii.PrintWithJustify(Words, WordsInArr, align, fileName, l, n, w)
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
