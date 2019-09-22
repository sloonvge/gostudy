package main
//
// import "fmt"
//
// var dict = map[byte]string{
// 	'2': "abc",
// 	'3': "def",
// 	'4': "ghi",
// 	'5': "jkl",
// 	'6': "mno",
// 	'7': "pqrs",
// 	'8': "tuv",
// 	'9': "wxyz",
// }
//
// func main() {
// 	number := ""
// 	fmt.Scan(&number)
// 	answer := make([]string, 0)
// 	generate(number, 0, make([]byte, 0), &answer)
// 	fmt.Printf("[")
// 	for i := 0; i < len(answer) - 1; i++ {
// 		fmt.Printf("%v, ", answer[i])
// 	}
// 	fmt.Printf("%s]", answer[len(answer) - 1])
// }
//
// func generate(num string, i int, b []byte, answer *[]string) {
// 	if i == len(num) {
// 		c := make([]byte, len(b))
// 		copy(c, b)
// 		*answer = append(*answer, string(c))
// 		return
// 	} else {
// 		s := dict[num[i]]
// 		for j := 0; j < len(s); j++ {
// 			b = append(b, s[j])
// 			generate(num, i + 1, b, answer)
// 			b = b[:len(b) - 1]
// 		}
// 	}
// }
