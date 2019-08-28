package offer

import (
	"math"
	"testing"
)

var (
	isInf = false
)

func Power(x float64, y int) float64 {
	if x == 0. && y < 0{
		isInf = true
		return math.Inf(1)
	}

	abs := uint(y)
	if y < 0 {
		abs = uint(-y)
	}
	ans := pow(x, abs)
	if y < 0 {
		return 1.0 / ans
	}
	return ans
}
func pow(x float64, y uint) float64 {
	result := 1.0
	for i := 0; uint(i) < y; i++ {
		result *= x
	}
	return result
}

func Power2(x float64, y int) float64 {
	if x == 0. && y < 0{
		isInf = true
		return math.Inf(1)
	}

	abs := uint(y)
	if y < 0 {
		abs = uint(-y)
	}
	ans := pow2(x, abs)
	if y < 0 {
		return 1.0 / ans
	}
	return ans
}
func pow2(x float64, y uint) float64 {
	if y == 0 {
		return 1.0
	}
	result := pow2(x, y >> 1)
	result *= result
	if y & 0x1 == 1 {
		result *= x
	}
	return result
}


func TestPower(t *testing.T) {
	inputs := []float64{
		0.0,
		1.0,
		-1.0,
		2.0,
		-2.0,
	}
	helper := []int{
		0, -1, 1, -2, 2,
	}
	wants := [][]float64{
		{1.0, math.Inf(1), 0.0, math.Inf(1), 0.0},
		{1.0, 1.0, 1.0, 1.0, 1.0},
		{1.0, -1.0, -1.0, 1.0, 1.0},
		{1.0, 0.5, 2.0, 0.25, 4.0},
		{1.0, -0.5, -2.0, 0.25, 4.0},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := Power(inputs[i],
					helper[j]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
	t.Run("2", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := math.Pow(inputs[i],
					float64(helper[j])); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
	t.Run("3", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := Power2(inputs[i],
					helper[j]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}

// 67
