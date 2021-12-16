package hashtable

import (
	"fmt"
	"testing"
)

// TestHashTable tests the HashTable struct
func TestHashTableValue(t *testing.T) {

	name := "Tonny"
	var expected int32 = 5

	h := Hash(name)
	fmt.Println(h)
	if h != expected {
		t.Error("expected", expected, "got", h)
	}
}

func TestHashTable(t *testing.T) {

	h := NewHashTable(10)
	err := h.Add("Tonny", "123")
	if err != nil {
		t.Error("expected", nil, "got", err)
	}

}