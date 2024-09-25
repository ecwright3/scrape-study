package main

import (
	"fmt"
	"os"
)

func main() {

	var url string

	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		fmt.Println("Enter Target URL:")
		fmt.Scan(&url)
	}

	getImages(url)
}
