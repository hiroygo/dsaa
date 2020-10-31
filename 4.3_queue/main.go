package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const processQueueMaxSize = 1024

type Process struct {
	Name      string
	UsageTime int
}

type ProcessQueue struct {
	pss   [processQueueMaxSize]Process
	begin int
	size  int
}

func (q *ProcessQueue) Size() int {
	return q.size
}

func (q *ProcessQueue) Enqueue(p Process) error {
	if q.Size() == processQueueMaxSize {
		return fmt.Errorf("queue overflow, max size is %v", processQueueMaxSize)
	}

	q.pss[(q.begin+q.size)%processQueueMaxSize] = p
	q.size++
	return nil
}

func (q *ProcessQueue) Dequeue() (Process, error) {
	if q.Size() == 0 {
		return Process{}, errors.New("queue underflow")
	}

	p := q.pss[q.begin]
	q.begin = (q.begin + 1) % processQueueMaxSize
	q.size--
	return p, nil
}

func executeProcessQueue(pq *ProcessQueue, quantum int) int {
	if pq == nil {
		return 0
	}

	for time := 0; pq.Size() != 0; {
		p, err := pq.Dequeue()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Dequeue error %v\n", err)
			return 1
		}

		min := math.Min(float64(quantum), float64(p.UsageTime))
		time += int(min)
		p.UsageTime -= int(min)
		if p.UsageTime <= 0 {
			fmt.Printf("%s %d\n", p.Name, time)
			continue
		}

		if err := pq.Enqueue(p); err != nil {
			fmt.Fprintf(os.Stderr, "Enqueue error %v\n", err)
			return 1
		}
	}
	return 0
}

func processQueueFromStdin() (*ProcessQueue, int, error) {
	sc := bufio.NewScanner(os.Stdin)

	// プロセス数とクオンタムの読み取り
	sc.Scan()
	heads := strings.Fields(sc.Text())
	if len(heads) != 2 {
		return nil, 0, errors.New("n q の形式で入力してください")
	}
	if err := sc.Err(); err != nil {
		return nil, 0, fmt.Errorf("Scanner error, %v", err)
	}
	psnum, err := strconv.Atoi(heads[0])
	if err != nil {
		return nil, 0, fmt.Errorf("n parse error, %v", err)
	}
	quantum, err := strconv.Atoi(heads[1])
	if err != nil {
		return nil, 0, fmt.Errorf("q parse error, %v", err)
	}

	// プロセス情報の読み取り
	pq := &ProcessQueue{}
	for i := 0; i < psnum; i++ {
		sc.Scan()
		fields := strings.Fields(sc.Text())
		if len(fields) != 2 {
			return nil, 0, errors.New("name usageTime の形式で入力してください")
		}

		p := Process{}
		p.Name = fields[0]
		if t, err := strconv.Atoi(fields[1]); err != nil {
			return nil, 0, fmt.Errorf("usageTime parse error, %v", err)
		} else {
			p.UsageTime = t
		}

		if err := pq.Enqueue(p); err != nil {
			return nil, 0, fmt.Errorf("Enqueue error, %v", err)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, 0, fmt.Errorf("Scanner error, %v", err)
	}

	return pq, quantum, nil
}

func main() {
	pq, quantum, err := processQueueFromStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Println("-----")
	os.Exit(executeProcessQueue(pq, quantum))
}
