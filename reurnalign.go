package ascii

import (
	"os"
	"strings"
)

func ReturnAlign() string {
	align := "left"
	for x := 0; x < len(os.Args); x++ {
		if strings.Contains(os.Args[x], "--align=") {
			align = strings.ToLower(strings.TrimPrefix(os.Args[x], "--align="))
		}
	}
	return align
}
