package main

import (
	"fmt"
	"testing"
	"time"
)

func abc(t *testing.T) {
	c := make(chan bool)
	startTimer(hello)
	<-c
}

func hello() {
	//每分钟执行一次
	fmt.Printf("hello timer.=== %v\n", time.Now().Local())
}

func startTimer(f func()) {
	go func() {
		for {
			now := time.Now()
			// 计算下一个时间点
			interval, _ := time.ParseDuration("1m") //
			fmt.Println(interval)
			next := now.Add(interval)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}
