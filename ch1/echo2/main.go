// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", i, arg)
	}
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start))

	start = time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start))
}

//!-
