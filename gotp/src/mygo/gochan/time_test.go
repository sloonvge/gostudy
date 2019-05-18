package main

import (
	"testing"
	"time"
	"fmt"
)

func TestTimeTicker(t *testing.T) {
	var peroid = time.Second * 3

	ticker := time.NewTicker(peroid)
	defer ticker.Stop()

	for   {
		<- ticker.C
		fmt.Println("ticker...")
	}
}
