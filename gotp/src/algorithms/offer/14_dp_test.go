package offer

import (
	"math"
	"testing"
)

// 14

func MaxProductAfterCutting1(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	a := maxProductAfterCutting1(n)
	return a
}
func maxProductAfterCutting1(n int) int {
	if n <= 3 {
		return n
	}
	a := 0
	for i := 1; i < n; i++ {
		a = intMax(a, maxProductAfterCutting1(i) * maxProductAfterCutting1(n - i))
	}
	return a
}

func MaxProductAfterCutting2(a int) (b int){
	if a < 2 {
		return 0
	}
	if a == 2 {
		return 1
	}
	if a == 3 {
		return 2
	}
	products := make([]int, a + 1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	var max int
	for i := 4; i <= a; i++ {
		max = 0
		for j := 1; j <= i / 2; j++ {
			max = intMax(max, products[j] * products[i - j])
		}
		products[i] = max
	}
	return products[a]

}

func MaxProductAfterCutting3(a int) (b int) {
	if a < 2 {
		return 0
	}
	if a == 2 {
		return 1
	}
	if a == 3 {
		return 2
	}

	timesOf3 := a / 3
	if a - timesOf3 *  3 == 1 {
		timesOf3 -= 1
	}
	timesOf2 := (a - timesOf3 * 3) / 2
	return int(math.Pow(3, float64(timesOf3)) * math.Pow(2, float64(timesOf2)))
}

func TestMaxProduct(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4, 8}
	wants := []int{0, 0, 1, 2, 4, 18}
	t.Run("1", func(t *testing.T) {
		for i, input := range inputs {
			got := MaxProductAfterCutting1(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	t.Run("2", func(t *testing.T) {
		for i, input := range inputs {
			got := MaxProductAfterCutting2(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	t.Run("3", func(t *testing.T) {
		for i, input := range inputs {
			got := MaxProductAfterCutting3(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}

func intMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
