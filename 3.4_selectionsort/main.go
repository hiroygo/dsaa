package main

import (
	"fmt"
	"os"

	"github.com/hiroygo/dsaa/common"
)

func selectionSort(ns []int) {
	for i := 0; i < len(ns); i++ {
		minj := i
		for j := i; j < len(ns); j++ {
			if ns[minj] > ns[j] {
				minj = j
			}
		}
		ns[i], ns[minj] = ns[minj], ns[i]
	}
}

func main() {
	ns, err := common.NumsFromArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	selectionSort(ns)

	fmt.Println(ns)
	os.Exit(0)
}
