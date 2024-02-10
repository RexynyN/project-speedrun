package solution

import (
	"fmt"
	"math"
	"slices"
)

func BinToDec(binary string) int {
	base, total := int(math.Pow(2, float64(len(binary)-1))), 0
	for _, run := range binary {
		if run == '1' {
			total += base
		}
		base /= 2
	}
	return total
}

func DecToBin(integer int) string {
	var binary []rune
	for integer > 0 {
		binary = append(binary, rune(fmt.Sprint(integer % 2)[0]))
		integer = integer / 2
	}
	slices.Reverse(binary)
	return string(binary)
}
