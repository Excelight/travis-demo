package main

import (
	"fmt"
	"github.com/Excelight/travis-demo/pkg"
)

func main() {
	res := pkg.Add(1, 2)
	fmt.Printf("res = %v\n", res)
}
