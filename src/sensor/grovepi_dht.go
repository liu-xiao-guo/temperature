package main

import (
	"fmt"
	"time"
)

func main() {
	var g GrovePi
	
	g = *InitGrovePi(0x07)
	
	for {
		time.Sleep(2 * time.Second)
		t, h, err := g.ReadDHT(D7)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
		fmt.Println(h)
	}
	
	g.CloseDevice()
}
