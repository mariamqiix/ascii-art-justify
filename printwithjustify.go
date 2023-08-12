package ascii

import (
	"fmt"
	"strconv"
)

func PrintWithJustify(Words [][]string, justify string, n, w, count int) string {
	align := 0
	if justify == "right" && n == 0 {
		return fmt.Sprintf("%"+strconv.Itoa(CheckPrint(Words, justify, count))+"v", Words[n][w])
	} else if justify == "center" && n == 0 {
		return fmt.Sprintf("%"+strconv.Itoa(CheckPrint(Words, justify, count)/2+3)+"v", Words[n][w])
	} else if justify == "justify" && n != 0 && Words[n-1][w] == "      " {
		if count == 0 {
			count = 1
		}
		return fmt.Sprintf("%"+strconv.Itoa((CheckPrint(Words, justify, count)/count)+3)+"v", Words[n][w])
	} else {
		return fmt.Sprintf("%"+strconv.Itoa(align)+"v", Words[n][w])

	}
	return ""
}
