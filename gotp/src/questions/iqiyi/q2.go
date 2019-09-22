package main

// 假设有N个人要玩游戏，每轮游戏必须选出一个人当裁判，剩下的N-1个人作为玩家。现在第i个人要求作为玩家进行至少Ai轮游戏，那么至少需要进行多少轮游戏才能满足所有人的要求？

// 第一行包含一个整数N，2≤N≤105。
//
// 第二行包含N个空格隔开的整数A1到AN，0≤Ai≤109。

// 3
// 2 2 3
// 4

// func main() {
// 	var n, m, k, l int
// 	fmt.Scan(&n)
// 	fmt.Scan(&m)
// 	numbersLen := int(math.Pow(2, float64(n)))
// 	numbers := make([]int, numbersLen)
// 	ops := make([][2]int, m)
// 	for i := 0; i < numbersLen; i++ {
// 		fmt.Scan(&k)
// 		numbers[i] = k
// 	}
// 	numbersCopy := make([]int, numbersLen)
// 	for i := 0; i < m; i++ {
// 		fmt.Scanf("%d %d", &k, &l)
// 		ops[i] = [2]int{k, l}
// 		numbers[k-1] = l
// 		copy(numbersCopy, numbers)
// 		num := operate(numbersCopy, n, numbersLen)
// 		fmt.Println(num)
// 	}
// }
//
// func operate(numbers []int, n, l int) int {
// 	answer := numbers[0]
// 	for _, n := range numbers[1:] {
// 		answer |= n
// 	}
//
// 	return answer
// }
