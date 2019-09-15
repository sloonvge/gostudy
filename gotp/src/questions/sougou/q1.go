package main
//
// import "fmt"
//
// type linkedList struct {
// 	value uint
// 	key string
// 	prev *linkedList
// 	next *linkedList
// }
//
// type DMap struct {
// 	dMap map[string]*linkedList
// 	size int
// 	cap int
// 	head *linkedList
// 	tail *linkedList
// }
//
// func NewDMap(n int) *DMap {
// 	return &DMap{
// 		dMap: make(map[string]*linkedList, 0),
// 		cap: n,
// 	}
// }
//
// func (m *DMap) Insert(key string, value uint) {
// 	old, ok := m.dMap[key]
// 	if !ok {
// 		l := &linkedList{value:value, key:key}
// 		if m.size < m.cap {
// 			m.dMap[key] = l
// 			m.size++
// 			m.addToChain(l)
// 		}else if m.size == m.cap {
// 			m.addToChain(l)
// 			l = m.tail
// 			if l != nil {
// 				m.removeFromChain(l)
// 				delete(m.dMap, l.key)
// 				fmt.Println(l.key, l.value)
// 			}
// 		}
// 	} else {
// 		if old.value < value {
// 			old.value = value
// 			m.removeFromChain(old)
// 			m.addToChain(old)
// 		}
// 	}
//
// }
//
// func (m *DMap) addToChain(l *linkedList) {
// 	l.prev = nil
// 	l.next = m.head
// 	if l.next != nil {
// 		l.next.prev = l
// 	}
// 	m.head = l
// 	if m.tail == nil {
// 		m.tail = l
// 	}
// }
//
// func (m *DMap) removeFromChain(l *linkedList) {
// 	if l == m.head {
// 		m.head = l.next
// 	}
// 	if l == m.tail {
// 		m.tail = l.prev
// 	}
// 	if l.next != nil {
// 		l.next.prev = l.prev
// 	}
// 	if l.prev != nil {
// 		l.prev.next = l.next
// 	}
// }
//
// func main() {
// 	var N int
// 	fmt.Scan(&N)
// 	var k string
// 	var v int
//
// 	dMap := NewDMap(N)
// 	for {
// 		n, _ := fmt.Scanf("%s %d", &k, &v)
// 		if n == 0 {
// 			break
// 		} else {
// 			dMap.Insert(k, uint(v))
// 		}
// 	}
// }
