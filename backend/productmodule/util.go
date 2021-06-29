package productmodule

import (
	"strconv"
	"strings"
)

func formatPrice(price int64) string {
	str := strconv.FormatInt(price, 10)
	result := ""
	parts := strings.Split(str, ".")
	for i, counter := len(parts[0])-1, 0; i >= 0; i, counter = i-1, counter+1 {
		if counter%3 == 0 && counter != 0 {
			result = "." + result
		}
		result = string((parts[0])[i]) + result
	}
	sign := ""
	if price < 0 {
		sign = "-"
	}
	if len(parts) > 1 {
		result = result + "," + parts[1]
	}

	return sign + "Rp" + result
}
