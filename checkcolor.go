package ascii

import (
	"strconv"
	"strings"
)

func CheckColor(userValue string) string {
	Colors := map[string]string{"red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m", "blue": "\033[34m", "purple": "\033[35m", "cyan": "\033[36m", "white": "\033[37m",
		"black": "\033[30m", "orange":"\033[38;5;208m"}
	for color, value := range Colors {
		if color == userValue {
			return value
		} else if strings.Index(userValue, "rgb") == 0 && userValue[len(userValue)-1] == ')' {
			c := (strings.Split(strings.TrimRight(strings.TrimPrefix(userValue, "rgb("), ")"), ","))
			for i := 0; i < len(c); i++ {
				a := strings.ReplaceAll(c[i], " ", "")
				check, err := strconv.Atoi(a)
				if err != nil || len(c) != 3 || check < 0 || check > 255 {
					return "NO"
				}
			}
			return userValue
		} else if strings.Index(userValue, "#") == 0 && len(userValue) == 7 {
			for i := 1; i <= 6; i++ {
				if (userValue[i] >= '0' && userValue[i] <= '9') || (userValue[i] >= 'a' && userValue[i] <= 'f') {
				} else {
					return "NO"
				}
			}
			return userValue
		}
	}
	return "NO"
}
