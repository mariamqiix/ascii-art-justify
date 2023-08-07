package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Validation() string {
	val := "yes"
	if len(os.Args) == 1 {
		Error()
	} else if strings.Index(os.Args[1], "--output=") == 0 {
		val = "output"
		if (len(os.Args) == 3 || len(os.Args) == 4) && len(os.Args[1]) > 9 && strings.Index(os.Args[1], ".txt") != -1 && len(os.Args[1]) == strings.Index(os.Args[1], ".txt")+4 {
			CheckLetter(os.Args[2])
			if len(os.Args) == 4 {
				CheckFont(os.Args[3])
			}
		} else {
			Error()
		}
		} else if strings.Index(os.Args[1], "--align=") == 0 {
			val = "justify"
			align := strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))
			if len(os.Args[1]) > 8 {
				if !(align == "justify" || align == "left" || align == "right" || align == "center"  ){
					Error()
				}
			} else {
				Error()
			}
	} else if strings.Index(os.Args[1], "--color=") == 0 {
		val = "color"
		color := strings.ToLower(strings.TrimPrefix(os.Args[1], "--color="))
		if len(os.Args[1]) > 9 {
			if CheckColor(color) == "NO" {
				Error()
			}
		} else {
			Error()
		}

		if len(os.Args) == 3 {
			CheckLetter(os.Args[2])
		} else if len(os.Args) == 4 {
			if os.Args[3] != "standard" && os.Args[3] != "shadow" && os.Args[3] != "thinkertoy" {
				val = "colorWletter"
				CheckLetter(os.Args[2])
				CheckLetter(os.Args[3])
				if strings.Index(os.Args[3], os.Args[2]) == -1 || len(os.Args[2]) == 0 {
					Error()
				}
			} else {

			}
		} else if len(os.Args) == 5 {
			if os.Args[4] == "standard" || os.Args[4] == "shadow" || os.Args[4] == "thinkertoy" {
				val = "colorWletterWfont"
				CheckLetter(os.Args[2])
				CheckLetter(os.Args[3])
				if strings.Index(os.Args[3], os.Args[2]) == -1 || len(os.Args[2]) == 0 {
					Error()
				}
			} else {
				Error()
			}
		} else if len(os.Args) == 6 {
			val = "colorW2letter"
			color2 := strings.ToLower(strings.TrimPrefix(os.Args[3], "--color="))
			if CheckColor(color2) == "NO"  { // || color2 == color
				Error()
			}
			if strings.Index(os.Args[3], "--color=") == 0 {
				CheckLetter(os.Args[2])
				CheckLetter(os.Args[4])
				CheckLetter(os.Args[5])
				if strings.Index(os.Args[5], os.Args[2]) == -1 || strings.Index(os.Args[5], os.Args[4]) == -1 || os.Args[2] == os.Args[4] {
					Error()
				}
			} else {
				Error()
			}
		} else {
			Error()
		}
	} else if len(os.Args) == 2 {
		CheckLetter(os.Args[1])
		val = "standard"
	} else if len(os.Args) == 3 {
		CheckLetter(os.Args[1])
		CheckFont(os.Args[2])
	} else {
		Error()
	}
	return val
}

func Error() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
	os.Exit(0)
}

func CheckLetter(s string) {
	for g := 0; g < len(s); g++ {
		if s[g] > 126 || s[g] < 32 {
			fmt.Println("ERROR: ascii letters only")
			os.Exit(0)
		}
	}
}

func CheckFont(s string) {
	FontType := strings.ToLower(s)
	if FontType != "standard" && FontType != "shadow" && FontType != "thinkertoy" {
		Error()
	}
}
