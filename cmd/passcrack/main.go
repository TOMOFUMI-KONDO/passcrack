package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TOMOFUMI-KONDO/passcrack"
)

func main() {
	lenFrom, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("failed to parse %s: %w", os.Args[1], err)
	}
	lenTo, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("failed to parse %s: %w", os.Args[2], err)
	}

	ans, err := passcrack.Run(lenFrom, lenTo)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Answer: %s\n", ans)
}
