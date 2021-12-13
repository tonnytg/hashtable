package hashtable

import (
	"fmt"
	"testing"
)

// TestHashTable tests the HashTable struct
func TestHashTable(t *testing.T) {

	name := "Tonny"
	var expected int32 = 5

	h := Hash(name)
	fmt.Println(h)
	if h != expected {
		t.Error("expected", expected, "got", h)
	}
}