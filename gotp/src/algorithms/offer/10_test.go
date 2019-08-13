package offer

import "testing"

func FibRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibRecursive(n - 1) + FibRecursive(n - 2)
}

func FibLoop(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	m := make([]int, n+1)
	m[0] = 0
	m[1] = 1
	for i := 2; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[n]
}

func FibLoop2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var a, b, ans int
	a = 0
	b = 1
	for i := 2; i <= n; i++ {
		ans = a + b
		a = b
		b = ans
	}

	return ans
}

func TestFib(t *testing.T) {
	t.Run("recursive", func(t *testing.T) {
		t.Log(FibRecursive(10))
	})
	t.Run("loop", func(t *testing.T) {
		t.Log(FibLoop(10))
	})
	t.Run("loop2", func(t *testing.T) {
		t.Log(FibLoop2(10))
	})
}
