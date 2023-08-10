package ascii

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintWithColor(Words [][]string, color, Text1, letter1, validation string) {
	colorB := color
	positions, colors := CheckPostion(Text1, colorB, letter1, validation)
	FlagB := false
	for w := 0; w < 8; w++ {
		index := 0
		stop := 0
		for n := 0; n < len(Words); n++ {
			if letter1 != "NO!!-" {
				if len(positions) != 0 {
					if positions[index] == n {
						FlagB = true
					}
					if FlagB {
						colorB = colors[index]
						stop++
						if (stop == len(letter1) && colorB == color) || ((validation == "colorW2L" || validation == "colorW2LF") && stop == len(os.Args[4]) && color != colors[index]) {
							FlagB = false
							stop = 0
							if index+1 < len(positions) {
								index++
							}
						}
					} else {
						colorB = "\033[0m"
					}
				} else {
					colorB = "\033[0m"
				}
			}
			if strings.Contains(colorB, "rgb") {
				colorB = strings.ReplaceAll(colorB, " ", "")
				c := (strings.Split(strings.TrimRight(strings.TrimPrefix(colorB, "rgb("), ")"), ","))
				red, _ := strconv.Atoi(c[0])
				green, _ := strconv.Atoi(c[1])
				blue, _ := strconv.Atoi(c[2])
				colorB = fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)
			} else if strings.Contains(colorB, "#") {
				colorB = colorB[1:]
				rgb, _ := strconv.ParseUint(colorB, 16, 32)
				colorB = fmt.Sprintf("\033[38;2;%d;%d;%dm", int(rgb>>16&0xFF), int(rgb>>8&0xFF), int(rgb&0xFF))
			}
			fmt.Print(colorB, Words[n][w])
		}
		if w+1 != 8 {
			fmt.Println()
		}
	}
	fmt.Println()
}
