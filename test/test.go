package main

import (
	"time"
	"fmt"
	"container/list"
)

func listTest()  {

}

func main()  {

	fmt.Printf("%v",time.Now().Unix())
	return
	l := list.New()
	//l.PushBack(1)
	////l.PushBack([]int{1,2,3})
	//l.PushBack(2)
	a:=l.Front()
	fmt.Printf("%s",a.Value)
	//l.Remove(a)
	//b:=l.Front()
	//fmt.Printf("%s",b.Value)
	//fmt.Printf("%s",l)
	return;
	ch := make(chan bool)
	c := make(chan int)
	//ticker := time.NewTicker(time.Second * 1)
	//t := time.After(time.Second*1)
	//go func() {
	//
	//}()
	////time.After(time.Second*10)
	//go func() {
	//	for _ = range t {
	//		fmt.Printf("ticked at %v", time.Now())
	//	}
	//}()
	timeout := time.After(time.Second * 2) //
	t1 := time.NewTimer(time.Second * 3)   // 效果相同 只执行一次
	var i int
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("channel sign")
				return
			case <-t1.C:      // 代码段2
				fmt.Println("3s定时任务")
			case <-timeout:   // 代码段1
				i++
				fmt.Println(i, "2s定时输出")
			case <-time.After(time.Second * 4):    // 代码段3
				fmt.Println("4s timeout。。。。")
			default:                               // 代码段4
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 6)
	close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")
	<- ch
}