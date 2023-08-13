package ascii

import "fmt"

func CheckPrint(s [][]string, justify string, count2 int) int {
	j := ""
	for n := 0; n < len(s); n++ {
			j += s[n][3]

	}
	width := width()-1
	if len(j) <= width {
		if len(j) == width {
			return 1
		}
		z := width - len(j)
		if justify != "justify" {
			return z
		}
		if z*count2+len(j) != width{
		if z*count2+len(j) > width{
			z = (z*count2)-((z*count2+len(j)-width))

			return  z
		} else if z*count2+len(j) < width {
			z = (z*count2)-((z*count2+len(j)-width))
			fmt.Print(z)
			return  z
		}
			} else {
				return z
			}
return z
		}
		
	
	return 1
}
