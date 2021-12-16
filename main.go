package main

import (
	"fmt"
	"github.com/tonnytg/hashtable/pkg/hashtable"
)

func main() {
	println("Start Hash Table")
	h := hashtable.NewHashTable(7)
	fmt.Println(h)
	//h.Add("key1", "value1")
	//h.Add("key2", "value2")
	//h.Add("key3", "value3")
	//
	//h.Get("Tonny")
	//
	//h.Remove("key2")
}
