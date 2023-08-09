package ascii

import (
	"fmt"
	"os"
	"strings"
)

func WriteFile(s [][]string, firstWord bool, validation string) {
	ArgsPass := os.Args
	fileName := strings.TrimPrefix(ArgsPass[1], "--output=")
	if validation == "outputWC2" || validation == "outputWCF2" {
		fileName = strings.TrimPrefix(ArgsPass[2], "--output=")
	} else if validation == "outputWCL2" || validation == "outputWCLF2" {
		fileName = strings.TrimPrefix(ArgsPass[3], "--output=")
	}
	if firstWord == true {
		os.Remove(fileName)
	}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error \n", err)
	} else {
		if len(s) != 0 {
			for w := 0; w < 8; w++ {
				for n := 0; n < len(s); n++ {
					file.WriteString(s[n][w])
				}
				if w+1 != 8 {
					file.WriteString("\n")
				}
			}
		}
		file.WriteString("\n")
	}
	file.Close()
}
