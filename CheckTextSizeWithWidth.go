package ascii

func CheckTextSizeWithWidth(s []string, filename string) bool {
	for _, m := range s {
		j := ""
		for _, c := range m {
			j += ReadLetter(byte(c), filename)[0]
		}
		width := width()
		if !(len(j) <= width) {
			return false
		}
	}
	return true
}
