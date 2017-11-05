package main

import (
	"time"
	"fmt"
)



func main()  {
	ch := make(chan bool)
	ticker := time.NewTicker(time.Second * 1)
	//time.After(time.Second*10)
	go func() {
		for _ = range ticker.C {
			fmt.Printf("ticked at %v", time.Now())
		}
	}()
	<- ch
}