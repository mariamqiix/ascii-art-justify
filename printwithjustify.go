package ascii

import (
	"fmt"
	"strconv"
)

func PrintWithJustify(Words [][] string,s [] string,justify,Text1 string, num int){
	for w := 0; w < 8; w++ {
		if len(Text1) == 0 {
			break
		}
		align := 0
		for n := 0; n < len(Words); n++ {
			if justify == "right" && n == 0 {
				align = CheckPrint(s, num)
			}
			if justify == "center" && n == 0 {
				fmt.Printf("%"+strconv.Itoa(CheckPrint(s, num)/2+3)+"v",Words[n][w] )
			} else {
				fmt.Printf("%"+strconv.Itoa(align)+"v", Words[n][w])

			}
			align = 0
		}
		if w+1 != 8 {
			fmt.Println()
		}
	}
	fmt.Println()
}