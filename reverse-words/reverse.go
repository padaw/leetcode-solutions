import (
	"math"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	length := len(words)
	mid := int(math.Floor(float64(length / 2)))
	for i := 0; i < mid; i++ {
		tmp := words[i]
		words[i] = words[length-i-1]
		words[length-i-1] = tmp
	}
	return strings.Join(words, " ")
}

