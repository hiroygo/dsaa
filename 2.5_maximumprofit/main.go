package main

import (
	"github.com/hiroygo/dsaa/common"

	"errors"
	"fmt"
	"os"
)

func submax(ns []int) (int, error) {
	if len(ns) < 2 {
		return 0, errors.New("要素数が不正です")
	}

	// 出現する数はすべて正の数なので (b - a) を計算する時
	// b より左側にある最小値を a に選べばいい
	min := ns[0]
	submax := ns[1] - min
	for i := 1; i < len(ns); i++ {
		n := ns[i]
		if val := n - min; submax < val {
			submax = val
		}
		// 次の submax 計算用に評価しておく
		if min > n {
			min = n
		}
	}

	return submax, nil
}

func main() {
	ns, err := common.NumsFromArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	ret, err := submax(ns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Println(ret)
	os.Exit(0)
}
