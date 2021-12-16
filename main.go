package main

import "fmt"

func Calc(s string) int {
	var sum int
	for _, v := range s {
		sum += int(v) - '0'
	}
	return sum % 7
}

type Cluster struct {
	ID int
	Node *Node
}

type Node struct {
	ID int
	NodeSlot *Slot
}

type Slot struct {
	ID int
	Value string
	NextSlot *Slot
}

var m = make([]Cluster, 7)

func init() {
	// Cria o slice de cluster


	// Alimenta o Cluster com os indices dos Nodes
	for i := 0; i < 7; i++ {
		m[i].ID = i
		m[i].Node = &Node{ID: i}
	}
}

func SaveData(v string) {
	if m[Calc(v)].Node.NodeSlot == nil {
		m[Calc(v)].Node.NodeSlot = &Slot{ID: Calc(v), Value: v}
	} else {
		m[Calc(v)].Node.NodeSlot.NextSlot = &Slot{ID: Calc(v), Value: v}
	}
}

func main() {

	names := []string{"Antonio Thomacelli Gomes", "Antonio Thomacelli Gomez"}

	for _, v := range names {
		// Cria o primeiro Slot e se estiver ocupado joga para o próximo
		SaveData(v)

		fmt.Printf("Valor: %s\nSalvo no Cluster: %v\t\t " +
			"Node:%v\t\t Slot: %v\n", v, Calc(v), m[Calc(v)].Node.ID, m[Calc(v)].Node.NodeSlot.ID)
		fmt.Println(m[Calc(v)].Node.NodeSlot)
		fmt.Println("---")
	}
}
