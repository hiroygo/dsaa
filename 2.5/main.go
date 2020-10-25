package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toIntSlice(ss []string) ([]int, error) {
	var ns []int
	for _, s := range ss {
		if n, err := strconv.Atoi(s); err != nil {
			return nil, fmt.Errorf("Atoi error, %s,", err)
		} else {
			ns = append(ns, n)
		}
	}
	return ns, nil
}

func numsFromStdin() ([]int, error) {
	var ss []string

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	for _, s := range strings.Fields(sc.Text()) {
		ss = append(ss, s)
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Scan error, %s,", err)
	}

	if ns, err := toIntSlice(ss); err != nil {
		return nil, fmt.Errorf("toIntSlice error, %s", err)
	} else {
		return ns, nil
	}
}

func numsFromArgs() ([]int, error) {
	var ss []string

	flag.Parse()
	for _, s := range flag.Args() {
		ss = append(ss, s)
	}

	if ns, err := toIntSlice(ss); err != nil {
		return nil, fmt.Errorf("toIntSlice error, %s", err)
	} else {
		return ns, nil
	}
}

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
	ns, err := numsFromArgs()
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
