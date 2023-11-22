package base62

import (
	"math"
	"slices"
	"strings"
)

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base10ToBase62(num uint64) string {
	result := []byte{}
	for num > 0 {
		div := num / 62

		mod := num % 62
		result = append(result, base[mod])
		num = div

	}
	slices.Reverse(result)
	return string(result)
}

func Base62StrToBase10(num string) uint64 {
	result := []byte(num)
	slices.Reverse(result)
	var sum uint64 = 0
	for i, v := range result {

		baseV := uint64(math.Pow(62, float64(i)))
		value := strings.Index(base, string(v))
		sum += baseV * uint64(value)

	}
	return sum
}
