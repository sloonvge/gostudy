package main

import "07Proxy/go/proxy"

func main() {
	mm := proxy.NewGirl("jiaojiao")
	daili := proxy.NewProxy(mm)
	daili.GiveChocolate()
	daili.GiveDolls()
	daili.GiveFlowers()
}
