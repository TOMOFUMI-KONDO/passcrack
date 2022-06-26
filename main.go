package passcrack

import (
	"fmt"
	"math"
)

var (
	nums   = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	lowers = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	uppers = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	// TODO: symbols = []string{}
)

func Run(lenFrom, lenTo int) (string, error) {
	chars := make([]rune, 0)
	chars = append(chars, nums...)
	chars = append(chars, lowers...)
	chars = append(chars, uppers...)

	var ans string

	for i := lenFrom; i <= lenTo; i++ {
		var j int64
		for j = 0; j < int64(math.Pow(float64(len(chars)), float64(i))); j++ {
			pass := gen(i, j, chars)

			res, err := send(pass)
			if err != nil {
				return "", fmt.Errorf("failed to send: %w", err)
			}

			if res {
				ans = pass
			}
		}
	}

	return ans, nil
}

func gen(length int, th int64, chars []rune) string {
	pass := make([]rune, 0, length)
	charLen := int64(len(chars))

	for i := 0; i < length; i++ {
		pass = append(pass, chars[th%charLen])
		th /= charLen
	}

	return string(pass)
}

func send(pass string) (bool, error) {
	return pass == "abcd", nil
}
