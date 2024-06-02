package main

import (
	"fmt"
	"time"

	flag "github.com/spf13/pflag"
)

var n *int = flag.IntP("number", "n", 1, "Number of documents to create")
var initialOffset *int = flag.IntP("offset", "o", 0, "Number of offset days (from today) to start documents from")

func init() {
	flag.Parse()
	if *n < 1 {
		*n = 1
	}
	if *initialOffset < 0 {
		*initialOffset = 0
	}
}

func getDateForOffset(offset int) string {
	o := offset + *initialOffset
	return time.Now().Add(time.Duration(o) * 24 * time.Hour).Format(time.DateOnly)
}

func main() {
	for i := 0; i < *n; i++ {
		day := getDateForOffset(i)
		fmt.Println(day)
	}
}
