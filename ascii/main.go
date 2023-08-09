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
	} else if validation == "colorWL" || validation == "colorWLF" || validation == "outputWC" || validation == "outputWCF" || validation == "outputWC2" || validation == "outputWCF2" {
		index += 2
	} else if validation == "colorW2L" || validation == "colorW2LF" {
		index += 4
	} else if validation == "outputWCL" || validation == "outputWCLF" || validation == "outputWCL2" || validation == "outputWCLF2" {
		index += 3
	}

	CheckLetter(os.Args[index])
	count := strings.Count(os.Args[index], " ")
	WordsInArr := strings.Split(os.Args[index], "\\n")

	fileName := "standard"
	if len(os.Args) == 3 && validation != "output" && validation != "color" && validation != "justify" {
		fileName = strings.ToLower(os.Args[index+1])
	} else if len(os.Args) == 4 && validation != "colorWL" && validation != "outputWC" && validation != "outputWC2" {
		fileName = strings.ToLower(os.Args[index+1])
	} else if validation == "colorWLF" || validation == "outputWCF" || validation == "outputWCF2" {
		fileName = strings.ToLower(os.Args[4])
	} else if validation == "outputWCLF" || validation == "outputWCLF2" {
		fileName = strings.ToLower(os.Args[5])
	} else if validation == "colorW2LF" {
		fileName = strings.ToLower(os.Args[6])
	}

	if ascii.CheckPrint(WordsInArr, -1, count, fileName) == -1 {
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
			ascii.WriteFile(Words, FirstWord, validation)
			FirstWord = false
		} else if validation == "color" || validation == "colorWL" || validation == "colorWLF" || validation == "colorW2L" || validation == "colorW2LF" {
			letter1 := "NO!!-"
			color := ascii.CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[1], "--color=")))
			if validation == "colorWL" || validation == "colorWLF" {
				letter1 = os.Args[index-1]
			} else if validation == "colorW2L" || validation == "colorW2LF" {
				letter1 = os.Args[2]
			}
			ascii.PrintWithColor(Words, color, Text1, letter1, validation, WordsInArr, fileName , l , count)
		} else {
			for w := 0; w < 8; w++ {
				if len(Text1) == 0 {
					break
				}
				for n := 0; n < len(Words); n++ {
					if validation == "justify" {
						align := strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))
						ascii.PrintWithJustify(Words, WordsInArr, align, fileName,"\033[0m", l, n, w, count)
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

func CheckLetter(s string) {
	for g := 0; g < len(s); g++ {
		if s[g] > 126 || s[g] < 32 {
			fmt.Println("ERROR: ascii letters only")
			os.Exit(0)
		}
	}
}
