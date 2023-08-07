package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	validation := ascii.Validation()
	index := 1
	if validation == "output" || validation == "color"  || validation == "justify" {
		index++
	} else if validation == "colorWletter" || validation == "colorWletterWfont" {
		index += 2
	} else if validation == "colorW2letter" {
		index += 4
	}

	WordsInArr := strings.Split(os.Args[index], "\\n")
	if ascii.CheckPrint(WordsInArr, -1) == -1 {
		fmt.Println("too much words , write less")
		return
	}
	fileName := "standard"
	if len(os.Args) == 3 && validation != "output" && validation != "color" && validation != "justify"  {
		fileName = strings.ToLower(os.Args[index+1])
	} else if len(os.Args) == 4 && validation != "colorWletter" {
		fileName = strings.ToLower(os.Args[index+1])
	} else if validation == "colorWletterWfont" {
		fileName = strings.ToLower(os.Args[4])
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
		if validation == "output" {
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
		} else if validation == "justify" {
			align := strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))
			ascii.PrintWithJustify(Words,WordsInArr,align,Text1,l)
		} else {
			for w := 0; w < 8; w++ {
				if len(Text1) == 0 {
					break
				}
				for n := 0; n < len(Words); n++ {
					fmt.Print(Words[n][w])
				}
				if w+1 != 8 {
					fmt.Println()
				}
			}
			fmt.Println()
		}
	}
}
