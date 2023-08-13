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
		OutputValdation(1)
		if len(os.Args) > 2 && strings.Index(os.Args[2], "--color=") == 0 {
			val = "outputWC"
			ColorValdation(2)
			if len(os.Args) > 3 && strings.Index(os.Args[3], "--align=") == 0 {
				val = "outputWCJ"
				JustifyValidation(3)
				if len(os.Args) == 6 {
					val = "outputWCJF"
				} else if len(os.Args) != 5 {
					Error()
				}
			} else if len(os.Args) > 4 && strings.Index(os.Args[4], "--align=") == 0 {
				val = "outputWCLJ"
				if len(os.Args) == 5 {
					Error()
				}
				LetterExistnce(5, 3)
				JustifyValidation(4)
				if len(os.Args) == 7 {
					val = "outputWCLJF"
				} else if len(os.Args) != 6 {
					Error()
				}
			} else if len(os.Args) == 5 {
				val = "outputWCL"
				if !CheckFont(os.Args[4]) {
					LetterExistnce(4, 3)
				} else {
					val = "outputWCF"
				}
			} else if len(os.Args) == 6 {
				val = "outputWCLF"
				LetterExistnce(4, 3)
			} else if len(os.Args) != 4 {
				Error()
			}
		} else if len(os.Args) > 2 && strings.Index(os.Args[2], "--align=") == 0 {
			val = "outputWJ"
			JustifyValidation(2)
			if len(os.Args) > 3 && strings.Index(os.Args[3], "--color=") == 0 {
				val = "outputWCJ"
				ColorValdation(3)
				if len(os.Args) == 6 {
					if CheckFont(os.Args[5]) {
						val = "outputWCJF"
					} else {
						val = "outputWCLJ"
						LetterExistnce(5, 4)
					}
				} else if len(os.Args) == 7 {
					val = "outputWCLJF"
					LetterExistnce(5, 4)
				} else if len(os.Args) != 5 {
					Error()
				}
			} else if len(os.Args) == 5 {
				val = "outputWJF"
			} else if len(os.Args) != 4 {
				Error()
			}
		} else if len(os.Args) != 3 && len(os.Args) != 4 {
			Error()
		}
	} else if strings.Index(os.Args[1], "--align=") == 0 {
		val = "Justify"
		JustifyValidation(1)
		if len(os.Args) > 2 && strings.Index(os.Args[2], "--output=") == 0 {
			val = "outputWJ"
			OutputValdation(2)
			if len(os.Args) > 3 && strings.Index(os.Args[3], "--color=") == 0 {
				ColorValdation(3)
				val = "outputWCJ"
				if len(os.Args) == 6 {
					if CheckFont(os.Args[5]) {
						val = "outputWCJF"
					} else {
						val = "outputWCLJ"
						LetterExistnce(5, 4)
					}
				} else if len(os.Args) == 7 {
					val = "outputWCLJF"
					LetterExistnce(5, 4)
				} else if len(os.Args) != 5 {
					Error()
				}
			} else if len(os.Args) == 5 {
				val = "outputWJF"
			} else if len(os.Args) != 4 {
				Error()
			}
		} else if len(os.Args) > 2 && strings.Index(os.Args[2], "--color=") == 0 {
			ColorValdation(2)
			if len(os.Args) > 3 && strings.Index(os.Args[3], "--output=") == 0 {
				OutputValdation(3)
				if len(os.Args) == 5 {
					val = "outputWCJ"
				} else if len(os.Args) == 6 {
					val = "outputWCJF"
				} else {
					Error()
				}
			} else if len(os.Args) > 4 && strings.Index(os.Args[4], "--output=") == 0 {
				val = "outputWCJL"
				OutputValdation(4)
				LetterExistnce(5, 3)
				if len(os.Args) == 7 {
					val = "outputWCJLF"
				} else if len(os.Args) != 6 {
					Error()
				}
			} else if len(os.Args) == 4 {
				val = "colorWJ"
			} else if len(os.Args) == 5 {
				if CheckFont(os.Args[4]) {
					val = "colorWJF"
				} else {
					val = "colorWJL"
					LetterExistnce(4, 3)
				}
			} else if len(os.Args) == 6 {
				val = "colorWJLF"
				LetterExistnce(4, 3)
			} else {
				Error()
			}
		} else if len(os.Args) != 3 && len(os.Args) != 4 {
			Error()
		}
	} else if strings.Index(os.Args[1], "--color=") == 0 {
		val = "color"
		ColorValdation(1)
		if len(os.Args) > 2 && strings.Index(os.Args[2], "--output=") == 0 {
			OutputValdation(2)
			val = "outputWC"
			if len(os.Args) > 3 && strings.Index(os.Args[3], "--align=") == 0 {
				JustifyValidation(3)
				val = "outputWCJ"
				if len(os.Args) == 6 {
					val = "outputWCJF"
				} else if len(os.Args) != 5 {
					Error()
				}
			} else if len(os.Args) == 5 {
				val = "outputWCF"
			} else if len(os.Args) != 4 {
				Error()
			}
		} else if len(os.Args) > 2 && strings.Index(os.Args[2], "--align=") == 0 {
			JustifyValidation(2)
			val = "colorWJ"
			if len(os.Args) > 3 && strings.Index(os.Args[3], "--output=") == 0 {
				OutputValdation(3)
				val = "outputWCJ"
				if len(os.Args) == 6 {
					val = "outputWCJF"
				} else if len(os.Args) != 5 {
					Error()
				}
			} else if len(os.Args) == 5 {
				val = "colorWJF"
			} else if len(os.Args) != 4 {
				Error()
			}
		} else if len(os.Args) > 3 && strings.Index(os.Args[3], "--align=") == 0 {
			JustifyValidation(3)
			val = "colorWLJ"
			if len(os.Args) > 4 && strings.Index(os.Args[4], "--output=") == 0 {
				OutputValdation(4)
				val = "outputWCLJ"
				if len(os.Args) == 7 {
					val = "outputWCLJF"
				} else if len(os.Args) != 6 {
					Error()
				}
				LetterExistnce(5, 2)
			} else if len(os.Args) == 6 {
				LetterExistnce(4, 2)
				val = "colorWLJF"
			} else if len(os.Args) == 5 {
				LetterExistnce(4, 2)
			} else {
				Error()
			}
		} else if len(os.Args) > 3 && strings.Index(os.Args[3], "--output=") == 0 {
			OutputValdation(3)
			val = "outputWLC"
			if len(os.Args) > 4 && strings.Index(os.Args[4], "--align=") == 0 {
				JustifyValidation(4)
				val = "outputWCLJ"
				if len(os.Args) == 7 {
					val = "outputWCLJF"
				} else if len(os.Args) != 6 {
					Error()
				}
				LetterExistnce(5, 2)
			} else if len(os.Args) == 6 {
				LetterExistnce(4, 2)
				val = "outputWLCF"
			} else if len(os.Args) == 5 {
				LetterExistnce(4, 2)
			} else {
				Error()
			}
		} else if len(os.Args) == 4 {
			val = "colorWF"
			if !CheckFont(os.Args[3]) {
				val = "colorWL"
				LetterExistnce(3, 2)
			}
		} else if len(os.Args) == 5 {
			val = "colorWLF"
			LetterExistnce(3, 2)
		} else if len(os.Args) == 6 || len(os.Args) == 7 {
			val = "colorW2L"
			if len(os.Args) == 7 {
				val = "colorW2LF"
			}
			if strings.Index(os.Args[3], "--color=") == 0 {
				ColorValdation(3)
				color := strings.ToLower(strings.TrimPrefix(os.Args[1], "--color="))
				color2 := strings.ToLower(strings.TrimPrefix(os.Args[3], "--color="))
				if color2 == color {
					fmt.Println("The colors should be different")
					Error()
				}
				LetterExistnce(5, 2)
				LetterExistnce(5, 4)
				if os.Args[2] == os.Args[4] || strings.IndexAny(os.Args[4], os.Args[2]) != -1 {
					fmt.Println("You must chose different letter") //edit the comment here
					Error()
				}
			} else {
				Error()
			}
		} else if len(os.Args) != 3 {
			Error()
		}

	} else if len(os.Args) != 2 && len(os.Args) != 3 {
		Error()
	}
	return val
}

func ColorValdation(i int) {
	if CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[i], "--color="))) == "NO" {
		fmt.Println("color does not exist")
		Error()
	}
}

func OutputValdation(i int) {
	if len(os.Args[i]) > 9 && strings.Index(os.Args[i], ".txt") != -1 && len(os.Args[i]) == strings.Index(os.Args[i], ".txt")+4 {
	} else {
		fmt.Println("The file name is incorrect")
		Error()
	}
}

func LetterExistnce(a, b int) {
	if strings.Index(os.Args[a], os.Args[b]) == -1 || len(os.Args[b]) == 0 {
		fmt.Println("letters to be colored does not exist in the word")
		Error()
	}
}

func JustifyValidation(i int) {
	align := strings.ToLower(strings.TrimPrefix(os.Args[i], "--align="))
	if !(align == "justify" || align == "left" || align == "right" || align == "center") {
		fmt.Println("align is incorrect, you can use (left, right, center, justify) only")
		Error()
	}
}

func CheckFont(s string) bool {
	FontType := strings.ToLower(s)
	if FontType != "standard" && FontType != "shadow" && FontType != "thinkertoy" {
		return false
	}
	return true
}

func Error() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("Example: go run . --align=right something standard")
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
