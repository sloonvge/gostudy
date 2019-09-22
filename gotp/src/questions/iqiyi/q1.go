package main

import (
	"fmt"
	"math"
)


// 首先给出一个长度为k=2^n（其中n为正整数）的序列A={a1，a2…a_{k-1},ak}，我们定义一个值v，这个值是由如下计算方法得到的，首先将A序列的第 i 位和第 i+1 位进行 OR 操作得到新数列A‘，然后再对A’序列操作，将A’ 序列的第 i 位和第 i+1 位进行 XOR 操作得到A‘’，对A‘’按照第一次操作方式进行OR操作，因为序列长度为2^n,所以显然进行n次操作之后就只剩下一个数字了，此时这个数字就是v。
//
// 例如A={1，2，3，4}，第一次操作之后为{1|2=3，3|4=7}，第二次操作后，{3^7=4},所以v=4。
//
// 但是显然事情并没有那么简单，给出A序列后，还有m个操作，每个操作表示为“a b”，表示将A[a]改为b，之后再对A序列求v值。（注意每一次对序列的修改的永久的，即第i次修改是建立在前i-1次修改之上的）

// 2 4
// 1 2 3 4
// 1 4
// 3 4
// 1 3
// 1 3

func main() {
	var n, m, k, l int
	fmt.Scan(&n)
	fmt.Scan(&m)
	numbersLen := int(math.Pow(2, float64(n)))
	numbers := make([]int, numbersLen)
	for i := 0; i < numbersLen; i++ {
		fmt.Scan(&k)
		numbers[i] = k
	}
	numbersCopy := make([]int, numbersLen)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &k, &l)
		numbers[k-1] = l
		copy(numbersCopy, numbers)
		num := operate(numbersCopy, n, numbersLen)
		fmt.Println(num)
	}
}

func operate(numbers []int, n, l int) int {
	var k int
	iter := 1
	for l > 1 {
		if iter % 2 == 1 {
			k = l / 2
			for i, j := 0, 0; i < l; i += 2 {
				a, b := numbers[i], numbers[i + 1]
				numbers[j] = a | b
				j++
			}
			numbers = numbers[:k]
			l = k
		}
		if iter % 2 == 0 {
			k = l / 2
			for i, j := 0, 0; i < l; i += 2 {
				a, b := numbers[i], numbers[i + 1]
				numbers[j] = a ^ b
				j++
			}
			numbers = numbers[:k]
			l = k
		}
		iter++
	}

	return numbers[0]
}
