package main

import "fmt"

func main() {
	var word string
	var step int
	fmt.Scan(&word)
	fmt.Scan(&step)


}

func deleteWord(bytes []byte, step int) {

}

func insert(b []byte, i int, s byte) bool {
	if i > len(b) {
		return false
	}
	b = append(b, s)
	b[i+1:] = b[i:len(b)-1]
	b[i] = s
	return true
}

func deletePart(b []byte, index int, s byte) {
	var i, j int
	i = index - 1
	j = index + 1
	for i >= 0 && j < len(b) {
		for b[i] == s && i >= 0{
			i--
		}
		for b[j] == s && j < len(b){
			j++
		}
		if j - i >= 2 {
			b[i:] = b[j + 1:]
			break
		}
	}
}
