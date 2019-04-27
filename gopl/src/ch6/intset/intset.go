package intset

import (
	"fmt"
	"strconv"
	"strings"
)

/** 
* Created by wanjx in 2019/4/14 22:08
**/
const BaseInt = 32 << (^uint(0) >> 63)
type IntSet struct {
	words []uint
}


func (s *IntSet) Has(x int) bool{
	word, bit := x/BaseInt, uint(x%BaseInt)

	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/BaseInt, uint(x%BaseInt)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	res := make([]string, 0)
	for i, word := range s.words {
		for j := 0; j < BaseInt; j++ {
			if word & (1 << uint(j)) != 0 {
				res = append(res, strconv.Itoa(BaseInt * i + j))
			}
		}

	}

	return strings.Join(res, " ")
}

// LianXi

func (s *IntSet) Len() int {
	if s == nil {
		return 0
	}

	var ret int
	for _, word := range s.words {
		n := 0
		for t := word; t != 0; t = t >> 1 {
			if  t & 1 == 1 {
				n++
			}
		}
		ret += n
	}
	return ret
}

func (s *IntSet) Remove(x int) {
	word, bit := x/BaseInt, uint(x%BaseInt)

	if s.Has(x) {
		s.words[word] &= ^(1 << bit)
	} else {
		fmt.Println("fail remove, not has value: ", x)
	}
}

func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

func (s *IntSet) Copy() *IntSet {
	var ret IntSet

	for _, word := range s.words {
		ret.words = append(ret.words, word)
	}

	return &ret
}

func (s *IntSet) AddAll(words ...int) {
	for _, word := range words {
		s.Add(word)
	}
}

// 6.3
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			continue
		}
		s.words[i] &= tword
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			continue
		}
		 s.words[i] = s.words[i] & ^(s.words[i] & tword)
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, tword)
		}
		s.words[i] = (s.words[i] | tword) & ^(s.words[i] & tword)
	}
}

// 6.4
func (s *IntSet) Elems() []int {
	res := make([]int, 0)
	for i, word := range s.words {
		for j := 0; j < BaseInt; j++ {
			if word & (1 << uint(j)) != 0 {
				res = append(res, BaseInt * i + j)
			}
		}
	}

	return res
}