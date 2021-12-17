package main

import "fmt"

type Node struct {
	ID int
	NodeSlot *Slot
}

type Slot struct {
	ID int
	Value string
	NextSlot *Slot
}

func NewSlot(s *Slot) *Slot{
	s = &Slot{}
	return s
}

func SaveData(value string, s *Slot) {
	fmt.Println(s)

	if s == nil {
		s = NewSlot(s)
		s.Value = value
		fmt.Println("true:", s)
	} else {
		SaveData(value, s.NextSlot)
		fmt.Println("false:", s)
		return
	}

	//if s == nil {
	//	fmt.Println("Slot is nil")
	//	s.Value = value
	//} else {
	//	if s.NextSlot == nil {
	//		s.NextSlot = NewSlot(s.NextSlot)
	//	}
	//	SaveData(value, s.NextSlot)
	//}
}

func main() {
	n := make([]Node, 3)
	n[0].ID = 0
	n[0].NodeSlot = &Slot{0, "a", nil}
	fmt.Printf("Node: %p\n", n)
	fmt.Printf("Slot: %p\n", n[0].NodeSlot)
	SaveData("Teste", n[0].NodeSlot)
	SaveData("Teste2", n[0].NodeSlot)


}