package test

import (
	"fmt"
	"testing"
)

// 1. How to define?
type Reader interface {
	Read() string
}

type Closer interface {
	Close() error
}

// 2. How to compose interfaces
type ReadCloser interface {
	Reader
	Closer
}

func readClose(rc ReadCloser) {
	fmt.Println(rc.Read())
	rc.Close()
}

type ReadCloserImp struct {
	data  string
	data2 string
}

func (r ReadCloserImp) Read() string {
	return r.data
}

func (r ReadCloserImp) Close() error {
	r.data = ""
	fmt.Println("Closing the obj")
	return nil
}

func TestInterfaces(t *testing.T) {
	rc := ReadCloserImp{"abc", "efg"}
	readClose(rc)
	fmt.Println(rc)
}
