package main

import "fmt"

func main() {
	test(1)

}

type any interface {
}

//go:noinline
func test[T1 any](y T1) {
	fmt.Printf("%v", y)
}
