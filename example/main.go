package main

import (
	"fmt"
	"github.com/wangjuelong/ringbuffer"
)

func main() {
	rb := ringbuffer.New(1024)
	rb.Write([]ringbuffer.ElementType{"a","b","c","d"})
	fmt.Println(rb.Length())
	fmt.Println(rb.Free())
	buf := make([]ringbuffer.ElementType, 4)

	rb.Read(buf)
	fmt.Println(buf)
	// Output: 4
	// 1020
	// [a,b,c,d]
}
