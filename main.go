package passcrack

import (
	"fmt"
	"log"
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

	for len_ := lenFrom; len_ <= lenTo; len_++ {
		var th uint64
		for th = 0; th < uint64(math.Pow(float64(len(chars)), float64(len_))); th++ {
			pass := genPass(len_, th, chars)

			res, err := crack(pass)
			if err != nil {
				return "", fmt.Errorf("failed to send: %w", err)
			}

			if res {
				return pass, nil
			}
		}
	}

	return "", fmt.Errorf("password was not found")
}

func RunConcurrent(lenFrom, lenTo int) (string, error) {
	chars := make([]rune, 0)
	chars = append(chars, nums...)
	chars = append(chars, lowers...)
	chars = append(chars, uppers...)

	ans := make(chan string, 1)

	for len_ := lenFrom; len_ <= lenTo; len_++ {
		var th uint64
		for th = 0; th < uint64(math.Pow(float64(len(chars)), float64(len_))); th++ {
			go func(i int, j uint64) {
				pass := genPass(i, j, chars)

				res, err := crack(pass)
				if err != nil {
					log.Printf("failed to send: %v", err)
					return
				}

				if res {
					ans <- pass
				}
			}(len_, th)
		}
	}

	return <-ans, nil
}

func genPass(length int, th uint64, chars []rune) string {
	pass := make([]rune, 0, length)
	charLen := uint64(len(chars))

	for i := 0; i < length; i++ {
		pass = append(pass, chars[th%charLen])
		th /= charLen
	}

	return string(pass)
}

func crack(pass string) (bool, error) {
	return pass == "ZZZZ", nil
}
