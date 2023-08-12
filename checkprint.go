package ascii

// import "fmt"

func CheckPrint(s [][]string, justify string, count2 int) int {
	j := ""
	for n := 0; n < len(s); n++ {
		if justify == "justify" {
			if s[n][4] != "      " {
				j += s[n][4]
			} else if s[n][4] == "      " && count2 < 4 {
				j += s[n][4]
			}
		} else {
			j += s[n][4]
		}
	}
	width := width()
	if len(j) <= width {
		if len(j) == width {
			return 1
		}
		z := width - len(j) + 4
		if justify != "justify" {
			return z
		}

		if z*count2 > width {
			if count2 > 4 {
				if z > width/(count2*4) {
					z -= count2*4 + 3
				}
			} else {
				z -= count2*3 + 3
			}

		}
		return z
	}
	return 1
}
