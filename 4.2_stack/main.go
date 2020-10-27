package main

import (
	"fmt"
	"os"

	"github.com/hiroygo/dsaa/common"
)

type intStack struct {
	ns []int
}

func (s *intStack) push(n int) {
	if s == nil {
		return
	}
	s.ns = append([]int{n}, s.ns...)
}

func (s *intStack) pop() int {
	if s == nil {
		return 0
	}

	if len(s.ns) == 0 {
		return 0
	}

	n := s.ns[0]
	s.ns = append(s.ns[:0], s.ns[1:]...)
	return n
}

func main() {
	_, err := common.NumsFromArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	stack := &intStack{}

	stack.push(10)
	stack.push(0)
	stack.push(2)
	fmt.Println(stack)

	stack.pop()
	fmt.Println(stack)

	os.Exit(0)
}
