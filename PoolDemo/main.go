package main

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
	"sync"
)

var bufioReaderPool sync.Pool

func newBufioReader(r io.Reader) *bufio.Reader {
	if v := bufioReaderPool.Get(); v != nil {
		br := v.(*bufio.Reader)
		br.Reset(r)
		return br
	}
	return bufio.NewReader(r)
}

func putBufioReader(br *bufio.Reader) {
	br.Reset(nil)
	bufioReaderPool.Put(br)
}

func mainn() {
	var copyBufPool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 32*1024)
			return &b
		},
	}

	bufp := copyBufPool.Get().(*[]byte)
	defer copyBufPool.Put(bufp)
}

func main()  {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b)

	p.Put(2)
	p.Put(3)
	a = p.Get().(int)
	runtime.GC() //手动调用GC
	b = p.Get().(int)
	fmt.Println(a, b)
}
