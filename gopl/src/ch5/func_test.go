package ch5

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
	"testing"
)


/**
* Created by wanjx in 2019/3/18 18:52
**/

/*
	函数的类型被称为函数的标识符，如果形参列表和
	返回值列表变量名一一对应，则函数具有相同的类型或标识符
*/

func TestRecursion(t *testing.T) {
	url := "https://golang.org"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		t.Logf("findlinks: %v\n", err)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
//
// func fetch(url string) io.Reader{
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 		os.Exit(1)
// 	}
// 	b, err := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 		os.Exit(1)
// 	}
// }

func visit(links []string, n *html.Node) []string{
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func TestRecursion2(t *testing.T) {
	url := "https://golang.org"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil{
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// 函数值使我们不仅可以通过数据来参数化函数，亦可以通过行为。
func TestFuncValue(t *testing.T) {
	t.Logf(strings.Map(func(r rune) rune {
		return r + 1
	}, "ABC"))
}

func TestFuncValue2(t *testing.T) {
	url := "https://golang.org"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
	}
	forEachNode(doc, startElement, endElement)
}
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
var depth int
func startElement(n *html.Node)  {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node)  {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		depth--
	}
}