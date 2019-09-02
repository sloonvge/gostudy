package base

import (
	"fmt"
	"testing"
)

/** 
* Created by wanjx in 2019/9/2 1:11
**/

func r1()  {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
}

func r2() {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)

}

func r3()  {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)

}

func TestRange(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1()
	})
	t.Run("2", func(t *testing.T) {
		r2()
	})
	t.Run("3", func(t *testing.T) {
		r3()
	})
}