package main

import (
	"fmt"
	"os"

	"github.com/hiroygo/dsaa/common"
)

func bubbleSort(ns []int) {
	continueSearch := true
	for continueSearch {
		continueSearch = false
		for i := len(ns) - 1; i >= 1; i-- {
			if ns[i-1] > ns[i] {
				ns[i-1], ns[i] = ns[i], ns[i-1]
				continueSearch = true
			}
		}
	}
}

func main() {
	ns, err := common.NumsFromArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	bubbleSort(ns)

	fmt.Println(ns)
	os.Exit(0)
}
