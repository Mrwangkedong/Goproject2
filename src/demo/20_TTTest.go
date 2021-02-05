package main

import "time"

type Human interface {
	Say() string
}

type Man struct {
}

func (m *Man) Say() string {
	return "man"
}

func IsNil(h interface{}) bool {
	return h == nil
}

func main() {

	ch := make(chan int)

	go func() {
		<-ch
	}()

	ch <- 1

	time.Sleep(1 * time.Second)
}
