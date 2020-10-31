package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdType int

const (
	unknown cmdType = iota
	insertX
	deleteX
	deleteFirst
	deleteLast
)

type cmd struct {
	tp  cmdType
	val int
}

func toCmdType(cmd string) (cmdType, error) {
	switch cmd {
	case "insert":
		{
			return insertX, nil
		}
	case "delete":
		{
			return deleteX, nil
		}
	case "deleteFirst":
		{
			return deleteFirst, nil
		}
	case "deleteLast":
		{
			return deleteLast, nil
		}
	default:
		{
			return unknown, fmt.Errorf("コマンド %s は不正です", cmd)
		}
	}
}

func cmdsFromStdin() ([]cmd, error) {
	sc := bufio.NewScanner(os.Stdin)

	// コマンド数の読み取り
	sc.Scan()
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("Scanner error, %v", err)
	}
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		return nil, fmt.Errorf("n parse error, %v", err)
	}

	// コマンドの読み取り
	var cmds []cmd
	for i := 0; i < n; i++ {
		sc.Scan()
		fields := strings.Fields(sc.Text())
		switch len(fields) {
		case 1:
			{
				t, err := toCmdType(fields[0])
				if err != nil {
					return nil, fmt.Errorf("toCmdType error, %v", err)
				}

				cmds = append(cmds, cmd{tp: t})
			}
		case 2:
			{
				t, err := toCmdType(fields[0])
				if err != nil {
					return nil, fmt.Errorf("toCmdType error, %v", err)
				}

				v, err := strconv.Atoi(fields[1])
				if err != nil {
					return nil, fmt.Errorf("Atoi error, %v", err)
				}

				cmds = append(cmds, cmd{tp: t, val: v})
			}
		default:
			{
				return nil, errors.New("<command> | <command> <value> の形式で入力してください")
			}
		}
	}

	return cmds, nil
}

type IntListNode struct {
	pPrev *IntListNode
	pNext *IntListNode
	val   int
}

type IntList struct {
	pHead *IntListNode
	size  int
}

func NewIntList() *IntList {
	return &IntList{pHead: &IntListNode{}}
}

func (list *IntList) Size() int {
	return list.size
}

func (list *IntList) deleteNode(p *IntListNode) {
	p.pPrev.pNext = p.pNext
	if p.pNext != nil {
		p.pNext.pPrev = p.pPrev
	}
	list.size--
}

func (list *IntList) getNode(idx int) *IntListNode {
	p := list.pHead.pNext
	for idx > 0 {
		p = p.pNext
		idx--
	}
	return p
}

// 連結リストの先頭に v を追加する
func (list *IntList) Insert(v int) {
	pNode := &IntListNode{pPrev: list.pHead, pNext: list.pHead.pNext, val: v}

	if list.pHead.pNext != nil {
		list.pHead.pNext.pPrev = pNode
	}
	list.pHead.pNext = pNode
	list.size++
}

// 連結リスト中の一番始めの v を削除する
func (list *IntList) Delete(v int) {
	if list.Size() == 0 {
		return
	}

	p := list.pHead.pNext
	for p != nil {
		if p.val == v {
			break
		}
		p = p.pNext
	}
	if p != nil {
		list.deleteNode(p)
	}
}

func (list *IntList) DeleteFisrt() {
	if list.Size() == 0 {
		return
	}

	p := list.getNode(0)
	list.deleteNode(p)
}

func (list *IntList) DeleteLast() {
	if list.Size() == 0 {
		return
	}

	p := list.getNode(list.Size() - 1)
	list.deleteNode(p)
}

func (list *IntList) At(idx int) (int, error) {
	if idx < 0 || idx >= list.Size() {
		return 0, fmt.Errorf("invalid index %d", idx)
	}

	p := list.getNode(idx)
	return p.val, nil
}

func runCmds(cmds []cmd) int {
	list := NewIntList()
	for i := range cmds {
		switch cmds[i].tp {
		case insertX:
			{
				list.Insert(cmds[i].val)
			}
		case deleteX:
			{
				list.Delete(cmds[i].val)
			}
		case deleteFirst:
			{
				list.DeleteFisrt()
			}
		case deleteLast:
			{
				list.DeleteLast()
			}
		}
	}

	// 結果
	for i := 0; i < list.Size(); i++ {
		if n, err := list.At(i); err != nil {
			fmt.Fprintf(os.Stderr, "At error, %v\n", err)
			return 1
		} else {
			fmt.Printf("%d ", n)
		}
	}
	fmt.Println()
	return 0
}

func main() {
	cmds, err := cmdsFromStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	os.Exit(runCmds(cmds))
}
