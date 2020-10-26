package common

import (
	"bufio"
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

func NumsFromStdin() ([]int, error) {
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

func NumsFromArgs() ([]int, error) {
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
