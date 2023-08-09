package ascii

func CheckPrint(s []string, i int, filename string) int {
	count := 0
	for _, m := range s {
		j := ""
		for _, c := range m {
			d := ReadLetter(byte(c), filename)[0]
			j += d
		}
		width := width()
		if len(j) <= width {
			if count == i {
				return width - len(j) + 4
			}
		} else {
			return -1
		}
		count++

	}
	return 1
}
