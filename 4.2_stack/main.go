package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intStack struct {
	ns []int
}

func (s *intStack) push(n int) error {
	if s == nil {
		return errors.New("nil error")
	}

	s.ns = append([]int{n}, s.ns...)
	return nil
}

func (s *intStack) pop() (int, error) {
	if s == nil {
		return 0, errors.New("nil error")
	}

	if size, err := s.size(); err != nil {
		return 0, fmt.Errorf("size error, %v", err)
	} else if size == 0 {
		return 0, errors.New("stack empty")
	}

	n := s.ns[0]
	s.ns = append(s.ns[:0], s.ns[1:]...)
	return n, nil
}

func (s *intStack) size() (int, error) {
	if s == nil {
		return 0, errors.New("nil error")
	}
	return len(s.ns), nil
}

func calc(tokens []string) (int, error) {
	stack := &intStack{}
	popTwoElems := func() (lhs, rhs int, err error) {
		if rhs, err = stack.pop(); err != nil {
			return 0, 0, err
		}
		if lhs, err = stack.pop(); err != nil {
			return 0, 0, err
		}
		return lhs, rhs, nil
	}

	for _, t := range tokens {
		// オペランドを push する
		if n, err := strconv.Atoi(t); err == nil {
			if err := stack.push(n); err != nil {
				return 0, fmt.Errorf("operand push error, %v", err)
			}
			continue
		}

		// オペレータを処理する
		result := 0
		switch t {
		case "*":
			{
				l, r, err := popTwoElems()
				if err != nil {
					return 0, fmt.Errorf("* popTwoElems error, %v", err)
				}
				result = l * r
			}
		case "+":
			{
				l, r, err := popTwoElems()
				if err != nil {
					return 0, fmt.Errorf("+ popTwoElems error, %v", err)
				}
				result = l + r
			}
		case "-":
			{
				l, r, err := popTwoElems()
				if err != nil {
					return 0, fmt.Errorf("- popTwoElems error, %v", err)
				}
				result = l - r
			}
		default:
			{
				return 0, fmt.Errorf("トークン %s は不正です", t)
			}
		}
		if err := stack.push(result); err != nil {
			return 0, fmt.Errorf("operator result push error, %v", err)
		}
	}

	// 最終的な計算結果
	if n, err := stack.pop(); err != nil {
		return 0, fmt.Errorf("result pop error, %v", err)
	} else {
		return n, nil
	}
}

func tokensFromStdin() ([]string, error) {
	var tokens []string

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	for _, s := range strings.Fields(sc.Text()) {
		tokens = append(tokens, s)
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Scanner error, %v", err)
	}

	return tokens, nil
}

func main() {
	tokens, err := tokensFromStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	n, err := calc(tokens)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Println(n)
	os.Exit(0)
}
