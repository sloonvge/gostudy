package offer

import (
	"fmt"
	"math"
	"testing"
)

// 60
func PrintProbability(number int) {
	if number < 1 {
		return
	}
	maxSum := number * 6
	probs := make([]int, maxSum - number + 1)
	computeProbability(number, probs)
	total := math.Pow(6, float64(number))
	for i := 0; i < len(probs); i++ {
		ratio := float64(probs[i]) / total
		fmt.Printf("%d: %e\n", i + number, ratio)
	}
}
func computeProbability(number int, probs []int) {
	for i := 1; i <= 6; i++ {
		probability(number, number, i, probs)
	}
}
func probability(original, current, sum int, probs []int) {
	if current  == 1 {
		probs[sum - original]++
	} else {
		for i := 1; i <= 6; i++ {
			probability(original, current - 1, i + sum, probs)
		}
	}
}

func PrintProbability2(number int) {
	if number < 1 {
		return
	}
	probs := make([][]int, 2)
	probs[0] = make([]int, 6 * number + 1)
	probs[1] = make([]int, 6 * number + 1)
	var flag int
	for i := 1; i <= 6; i++ {
		probs[flag][i] = 1
	}
	for k := 2; k <= number; k++ {
		for i := 0; i < k; i++ {
			probs[1 - flag][i] = 0
		}
		for i := k; i <= 6 * k; i++ {
			probs[1 - flag][i] = 0
			for j := 1; j <= i && j <= 6; j++ {
				probs[1-flag][i] += probs[flag][i-j]
			}
		}
		flag = 1 - flag
	}
	total := math.Pow(6, float64(number))
	for i := number; i <= 6 * number; i++ {
		ratio := float64(probs[flag][i]) / total
		fmt.Printf("%d: %e\n", i, ratio)
	}
}

func TestPrintProbability(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		PrintProbability(2)
	})
	t.Run("2", func(t *testing.T) {
		PrintProbability2(2)
	})
}
