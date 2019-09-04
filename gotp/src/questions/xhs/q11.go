package main
//
// import "fmt"
//
// func main() {
// 	var s string
// 	fmt.Scan(s)
// 	printString(s)
// }
//
// func printString(s string) string {
// 	s = deletePar(s)
// 	s = deleteLeft(s)
// 	return s
// }
// func deletePar(s string) string {
// 	t := 0
// 	start := 0
// 	end := 0
// 	indexs := make([]int, 0)
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '(' {
// 			t++
// 			if t == 1 {
// 				start = i
// 			}
// 		}
// 		if s[i] == ')' {
// 			t--
// 			if t == 0 {
// 				end = i
// 				indexs = append(indexs, start)
// 				indexs = append(indexs, end)
// 			}
// 		}
//
// 	}
// 	var ret []byte
// 	begin := 0
// 	for i := 0; i < len(indexs); i += 2 {
// 		start = indexs[i]
// 		end = indexs[i+1]
// 		ret = append(ret, s[begin:start]...)
// 		begin = end + 1
// 	}
// 	return string(ret)
// }
//
// func deleteLeft(s string) string{
// 	lefts := make([]int, 0)
// 	for i := 0; i < len(s); i++ {
// 		if i > 1 {
// 			if s[i] == '<' && s[i-1] != '<'{
// 				lefts = append(lefts, i)
// 			}
// 		}
// 	}
// 	var ret []byte
// 	for i := 0; i < len(s); i++ {
// 		if i > 1 {
// 			if inLefts(lefts, i-1) && s[i] == '<'{
// 				continue
// 			}
// 		}
// 		if s[i] == '<' {
// 			continue
// 		}
// 		ret = append(ret, s[i])
//
// 	}
// 	return string(ret)
// }
//
// func inLefts(ls []int, i int) bool {
// 	for _, l := range ls {
// 		if i == l {
// 			return true
// 		}
// 	}
// 	return false
// }
