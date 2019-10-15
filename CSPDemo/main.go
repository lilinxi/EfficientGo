package main

import (
	"fmt"
	"time"
)

type Cmd struct {
	A        int
	B        int
	Sum      int
	doneChan chan *Cmd
}

type Calculator struct {
	cQueue chan *Cmd
}

func (cl *Calculator) calculate() {
	for c := range cl.cQueue {
		c.Sum = c.A + c.B
		c.doneChan <- c
	}
	fmt.Println("end")
}

func (cl *Calculator) sumAsync(a, b int) *Cmd {
	c := &Cmd{A: a, B: b, doneChan: make(chan *Cmd)}
	cl.cQueue <- c
	return c
}

func (cl *Calculator) Sum(a, b int) int {
	c := cl.sumAsync(a, b)
	res := <-c.doneChan
	return res.Sum
}

func NewCalculator() *Calculator {
	cl := &Calculator{cQueue: make(chan *Cmd, 10)}
	go cl.calculate()
	return cl
}

func main() {
	cl := NewCalculator()
	sum := cl.Sum(2, 3)
	time.Sleep(time.Second*2)
	fmt.Println(sum)
}
