# ringbuffer

[![License](https://img.shields.io/:license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/smallnest/ringbuffer?status.png)](http://godoc.org/github.com/smallnest/ringbuffer)  [![travis](https://travis-ci.org/smallnest/ringbuffer.svg?branch=master)](https://travis-ci.org/smallnest/ringbuffer) [![Go Report Card](https://goreportcard.com/badge/github.com/smallnest/ringbuffer)](https://goreportcard.com/report/github.com/smallnest/ringbuffer) [![coveralls](https://coveralls.io/repos/smallnest/ringbuffer/badge.svg?branch=master&service=github)](https://coveralls.io/github/smallnest/ringbuffer?branch=master) 

A circular buffer (ring buffer) in Go, implemented io.ReaderWriter interface

[![wikipedia](Circular_Buffer_Animation.gif)](https://github.com/smallnest/ringbuffer)


```go
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

```
