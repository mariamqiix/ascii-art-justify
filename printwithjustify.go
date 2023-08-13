package ascii

import (
	"fmt"
	"strings"
)

func PrintWithJustify(Words [][]string, justify string, n, w, count int) string {
	align := 0
	if justify == "right" && n == 0 {
		align = CheckPrint(Words, justify, count)
		return fmt.Sprintf(strings.Repeat(" ",align) + Words[n][w])
	} else if justify == "center" && n == 0 {
		align = CheckPrint(Words, justify, count)/2
		return fmt.Sprintf(strings.Repeat(" ",align) + Words[n][w])
	} else if justify == "justify" && n != 0 && Words[n][3] == "      " &&  Words[n][5] == "      " && Words[n][7] == "      " {
		if count == 0 {
			count = 1
		}
		align = CheckPrint(Words, justify, count)/count
		return fmt.Sprintf(strings.Repeat(" ",align)+Words[n][w])
	} else {
		return fmt.Sprintf( Words[n][w])

	}
	return ""
}
