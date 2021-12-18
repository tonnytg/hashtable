package main

import (
	"github.com/tonnytg/hashtable/pkg/hashtable"
)

func main() {

	hashTable := hashtable.Init()
	names := []string {
		"Antonio Thomacelli Gomes",
		"Antonio Thomacelli Gomez",
		"Antonio Thomacelli Gones",
	}

	for _, v := range names {
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
	hashTable.Search("STAN")
	hashTable.Insert("STAN")
	hashTable.Search("STAN")

}
