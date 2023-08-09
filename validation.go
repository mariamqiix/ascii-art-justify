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
		if len(os.Args[1]) > 9 && strings.Index(os.Args[1], ".txt") != -1 && len(os.Args[1]) == strings.Index(os.Args[1], ".txt")+4 {
			if len(os.Args) < 3 {
				Error()
			} else if strings.Index(os.Args[2], "--color=") == 0 {
				val = "outputWC"
				if len(os.Args[2]) > 9 {
					color := strings.ToLower(strings.TrimPrefix(os.Args[2], "--color="))
					if CheckColor(color) == "NO" {
						fmt.Println("color does not exist")
						Error()
					}
				} else {
					Error()
				}
				if len(os.Args) == 4 {
				} else if len(os.Args) == 5 {
					val = "outputWCL"
					if !CheckFont2(os.Args[4]) {
						if strings.Index(os.Args[4], os.Args[3]) == -1 || len(os.Args[3]) == 0 {
							Error()
						}
					} else {
						val = "outputWCF"
						CheckFont(os.Args[4])
					}
				} else if len(os.Args) == 6 {
					val = "outputWCLF"
					CheckFont(os.Args[5])
					if strings.Index(os.Args[4], os.Args[3]) == -1 || len(os.Args[3]) == 0 {
						Error()
					}
				} else {
					Error()
				}
			} else if strings.Index(os.Args[1], "--align=") == 0 { //do here

			} else if len(os.Args) == 3 || len(os.Args) == 4 {
				if len(os.Args) == 4 {
					CheckFont(os.Args[3])
				}
			} else {
				Error()
			}
		} else {
			Error()
		}

	} else if strings.Index(os.Args[1], "--align=") == 0 {

		val = "justify"
		align := strings.ToLower(strings.TrimPrefix(os.Args[1], "--align="))
		if len(os.Args) == 3 || len(os.Args) == 4 {
			if len(os.Args[1]) > 8 {
				if !(align == "justify" || align == "left" || align == "right" || align == "center") {
					Error()
				}
				if len(os.Args) == 4 {
					CheckFont(os.Args[3])
					val = "justify"
				}
			} else {
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
				fmt.Println("color does not exist")
				Error()
			}
		} else {
			Error()
		}
		if len(os.Args) == 3 {
		} else if len(os.Args) == 4 {
			if strings.Index(os.Args[2], "--output=") == 0 {
				val = "outputWC2"
				if len(os.Args[2]) > 9 && strings.Index(os.Args[2], ".txt") != -1 && len(os.Args[2]) == strings.Index(os.Args[2], ".txt")+4 {
				} else {
					Error()
				}
			} else if !CheckFont2(os.Args[3]) {
				val = "colorWL"
				if strings.Index(os.Args[3], os.Args[2]) == -1 || len(os.Args[2]) == 0 {
					fmt.Println("letters to be colored does not exist")
					Error()
				}
			}
		} else if len(os.Args) == 5 {
			if CheckFont2(os.Args[4]) {
				if strings.Index(os.Args[2], "--output=") == 0 {
					val = "outputWCF2"
				} else {
					val = "colorWLF"
					if strings.Index(os.Args[3], os.Args[2]) == -1 || len(os.Args[2]) == 0 {
						fmt.Println("letters to be colored does not exist")
						Error()
					}
				}
			} else if strings.Index(os.Args[3], "--output=") == 0 {
				val = "outputWCL2"
				if strings.Index(os.Args[4], os.Args[2]) == -1 || len(os.Args[2]) == 0 {
					fmt.Println("letters to be colored does not exist")
					Error()
				}
			} else {
				Error()
			}
		} else if len(os.Args) == 6 || len(os.Args) == 7 {
			val = "colorW2L"
			if len(os.Args) == 7 {
				val = "colorW2LF"
				CheckFont(os.Args[6])
			}
			if strings.Index(os.Args[3], "--color=") == 0 {
				color2 := strings.ToLower(strings.TrimPrefix(os.Args[3], "--color="))
				if CheckColor(color2) == "NO" || color2 == color {
					fmt.Println("The colors should be different")
					Error()
				}
				if strings.Index(os.Args[5], os.Args[2]) == -1 || strings.Index(os.Args[5], os.Args[4]) == -1 || os.Args[2] == os.Args[4] || len(os.Args[2]) == 0 || len(os.Args[4]) == 0 {
					fmt.Println("letters to be colored does not exist")
					Error()
				}
			} else if strings.Index(os.Args[3], "--output=") == 0 && len(os.Args) == 6 {
				val = "outputWCLF2"
				if strings.Index(os.Args[4], os.Args[2]) == -1 || len(os.Args[2]) == 0 {
					fmt.Println("letters to be colored does not exist")
					Error()
				}
				CheckFont(os.Args[5])
			} else {
				Error()
			}
		} else {
			Error()
		}
	} else if len(os.Args) == 2 {
	} else if len(os.Args) == 3 {
		CheckFont(os.Args[2])
	} else {
		Error()
	}
	return val
}

func Error() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("Example: go run . --align=right something standard")
	os.Exit(0)
}

func CheckFont(s string) {
	FontType := strings.ToLower(s)
	if FontType != "standard" && FontType != "shadow" && FontType != "thinkertoy" {
		Error()
	}
}

func CheckFont2(s string) bool {
	FontType := strings.ToLower(s)
	if FontType != "standard" && FontType != "shadow" && FontType != "thinkertoy" {
		return false
	}
	return true
}
