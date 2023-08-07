package ascii

import "os"


func CheckPrint(s []string, i int) int {
	count := 0
	for _, m := range s {
		j := ""
		for _, c := range m {
			file := "standard"
			if len(os.Args) == 4 {
				file = os.Args[3]
			}
			d := ReadLetter(byte(c), file)[0]
			j += d
		}
		width := width()
		if len(j) <= width {
			if count == i {
				return width - len(j)+4
			}
		} else {
			return -1
		}
		count++

	}
	return 1
}
