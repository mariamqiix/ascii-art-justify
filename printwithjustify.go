package ascii

import (
	"fmt"
	"strconv"
)

func PrintWithJustify(Words [][]string, s []string, justify, fileName string, num, n, w int) {
	align := 0
	if justify == "right" && n == 0 {
		align = CheckPrint(s, num, fileName)
	}
	if justify == "center" && n == 0 {
		fmt.Printf("%"+strconv.Itoa(CheckPrint(s, num, fileName)/2+3)+"v", Words[n][w])
	} else {
		fmt.Printf("%"+strconv.Itoa(align)+"v", Words[n][w])

	}
}
