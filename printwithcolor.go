package ascii

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintWithColor(Words [][]string, color, Text1, letter1, validation string) {
	colorB := color
	nbr, colors := CheckPostion(Text1, colorB, letter1, validation)
	FlagB := false
	for w := 0; w < 8; w++ {
		if len(Text1) == 0 {
			break
		}
		d := 0
		a := 0
		for n := 0; n < len(Words); n++ {
			if letter1 != "NO!!-" {
				if len(nbr) != 0 {
					if nbr[d] == n {
						FlagB = true
					}
					if FlagB {
						colorB = colors[d]
						a++
						if (a == len(letter1) || (validation == "colorW2letter" && a == len(os.Args[4]))) && colorB == colors[d] {
							FlagB = false
							a = 0
							if d+1 < len(nbr) {
								d++
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
				var a []int
				for i := 0; i < len(c); i++ {
					check, _ := strconv.Atoi(c[i])
					a = append(a, check)
				}
				red := a[0]
				green := a[1]
				blue := a[2]
				colorB = fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)
			} else if strings.Contains(colorB, "#") {
				colorB = colorB[1:]
				rgb, _ := strconv.ParseUint(colorB, 16, 32)
				red := int(rgb >> 16 & 0xFF)
				green := int(rgb >> 8 & 0xFF)
				blue := int(rgb & 0xFF)
				colorB = fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)
			}
			fmt.Print(colorB, Words[n][w])
		}
		if w+1 != 8 {
			fmt.Println()
		}
	}
	fmt.Println()
}
