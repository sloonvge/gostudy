package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ip := ""
	fmt.Scan(&ip)
	isIPv4 := strings.Contains(ip, ".") && ipv4(ip)

	isIPv6 :=  strings.Contains(ip, ":") && ipv6(ip)
	if !isIPv4 && !isIPv6 {
		fmt.Println("Neither")
	} else if isIPv4 {
		fmt.Println("IPv4")
	} else if isIPv6 {
		fmt.Println("IPv6")
	}
}

func ipv4(ip string) bool {
	isValid := true
	strs := strings.Split(ip, ".")
	valid := func(s string) bool{
		n, err := strconv.Atoi(s)
		if err == nil {
			return false
		}
		if n < 0 || n > 255 {
			return false
		}
		if len(s) > 1 && len(s) < 4 && strings.HasPrefix(s, "0"){
			return false
		}

		return true
	}
	for _, str := range strs {
		if !valid(str) {
			isValid = false
			break
		}
	}

	return isValid
}


func ipv6(ip string) bool {
	return false
}

