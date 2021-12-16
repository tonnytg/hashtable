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

func main() {

	// Cria o slice de cluster
	m := make([]Cluster, 7)

	// Alimenta o Cluster com os indices dos Nodes
	for i := 0; i < 7; i++ {
		m[i].ID = i
		m[i].Node = &Node{ID: i}
	}

	// Percorre o Cluster e imprime o Slot somente se tiver conteúdo
	for i := 0; i < 7; i++ {
		fmt.Println(m[i])
		if m[i].Node.NodeSlot != nil {
			fmt.Println(m[i].Node.NodeSlot)
		}
	}

	nome := "Antonio Thomacelli Gomes"
	fmt.Println("Será armazenado em:", Calc(nome))

	// Cria o primeiro Slot e se estiver ocupado joga para o próximo
    if m[Calc(nome)].Node.NodeSlot == nil {
		m[Calc(nome)].Node.NodeSlot = &Slot{ID: Calc(nome), Value: nome}
	} else {
		m[Calc(nome)].Node.NodeSlot.NextSlot = &Slot{ID: Calc(nome), Value: nome}
	}

	fmt.Println(m[Calc(nome)].Node.NodeSlot)

	//fmt.Println(m)
	//c := Calc(m[1])
	//fmt.Println(c)

	//println("Start Hash Table")
	//h := hashtable.NewHashTable(7)
	//fmt.Println(h)
	//h.Add("key1", "value1")
	//h.Add("key2", "value2")
	//h.Add("key3", "value3")
	//
	//h.Get("Tonny")
	//
	//h.Remove("key2")
}
