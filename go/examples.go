package main

import (
	"fmt"
)

type mytest struct {
	m map[string]int
}

func (t *mytest) getMap() map[string]int {
	//t.m["hello"] = 1
	r := make(map[string]int)
	for k, v := range t.m {
		r[k] = v
	}
	return r 
}

func main() {
	t := mytest{
		m: map[string]int{
			"hello":1,
		},
	}

	m := t.getMap()
	fmt.Println(m)
	m["world"] = 2
	fmt.Println(m)
	fmt.Println(t.m)
}