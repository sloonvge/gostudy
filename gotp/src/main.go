package main

import (
	"fmt"
	"io"
)

/**
* Created by wanjx in 2019/4/5 20:21
**/

func init() {
	n := 0
	s := ""
	var ans int
	fmt.Scan(&n)
	fmt.Scan(&s)
	q1 := &queue{
		Q: make([]string, 0),
	}
	q0 := &queue{
		Q: make([]string, 0),
	}
	
	for _, x := range s {
		if string(x) == "1" {
			if len(q0.Q) == 0 {
				q1.Push(string(x))
			} else {
				q0.Pop()
			}
		}
		if string(x) == "0" {
			if len(q1.Q) == 0 {
				q0.Push(string(x))
			} else {
				q1.Pop()
			}
		}
	}

	ans = len(q1.Q) + len(q0.Q)
	fmt.Printf("%d\n", ans)
}

type queue struct {
	Q []string
}

func (q *queue) Push(x string) {
	q.Q = append(q.Q, x)
}

func (q *queue) Pop() string {
	n := len(q.Q)
	if n == 0 {
		return "-1"
	}
	v := q.Q[n - 1]
	q.Q = q.Q[:n - 1]
	return v
}

func main() {
	// input := make([][]int, 0)
	var a string
	for {
		_, err := fmt.Scanf("%s", &a)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			fmt.Printf("%q %d\n",a, len(a))
		}
	}
}