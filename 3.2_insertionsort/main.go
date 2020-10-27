package main

import (
	"fmt"
	"os"

	"github.com/hiroygo/dsaa/common"
)

func insertionSort(ns []int) {
	for i := 1; i < len(ns); i++ {
		// 挿入する値を保存しておく
		n := ns[i]

		// 挿入するときはすでにソート済の範囲をずらす
		j := i - 1
		for ; j >= 0 && ns[j] > n; j-- {
			ns[j+1] = ns[j]
		}

		// j の for を抜けると、挿入したいインデックスより -1 されているので +1 する
		ns[j+1] = n
	}
}

func main() {
	ns, err := common.NumsFromArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	insertionSort(ns)

	fmt.Println(ns)
	os.Exit(0)
}
