package SinglyLinkedList

import (
	"fmt"
	"math/rand"
)

type MyNode struct {
	Data    int
	Pointer *MyNode
}

type MyList struct {
	Head MyNode
}

func CreateNode(data int, node *MyNode) MyNode {
	nextNode := MyNode{data,
		node}

	return nextNode
}

func CrateLinkedList(count int) MyList {
	var linkedList MyList
	var pointer *MyNode = nil

	for i := count; i > 0; i-- {
		data := rand.Intn(100)
		node := CreateNode(data, pointer)

		if i == 1 {
			linkedList = MyList{node}
		} else {
			pointer = &node
		}
	}

	return linkedList
}

func PrintLinkedList(linkedList MyList) {
	var nextNode MyNode

	pointer := linkedList.Head.Pointer
	fmt.Printf("Data: %d, Pointer: %p\n", linkedList.Head.Data, pointer)

	for {
		nextNode = *pointer
		pointer = nextNode.Pointer
		fmt.Printf("Data: %d, Pointer: %p\n", nextNode.Data, pointer)

		if pointer == nil {
			break
		}
	}

	fmt.Printf("%v", "\n")
}

func SearchItem(linkedList MyList, data int) MyNode {
	node := &(linkedList.Head)

	for node.Pointer != nil && node.Data != data {
		node = node.Pointer
	}

	return *node
}

func InsertItemFront(linkedList MyList, data int) MyList {
	// Создание узла
	newNode := CreateNode(data, &linkedList.Head)
	// Установка head на новый узел
	newList := MyList{Head: newNode}

	return newList
}

func InsertItemTail(linkedList MyList, data int) MyList {
	node := &(linkedList.Head)

	// Выбор узла с нулевым указателем
	for node.Pointer != nil {
		node = node.Pointer
	}

	// Создание узла
	newNode := CreateNode(data, nil)
	// Установка указателя предпоследнего узла на новый узел
	(*node).Pointer = &newNode

	return linkedList
}

func InsertItemAfterSelectNode(linkedList MyList, index int, data int) MyList {
	node := &(linkedList.Head)
	counter := 0

	// Итерация до искомого узла
	for counter < index && node.Pointer != nil {
		node = node.Pointer // &1
		counter++
	}

	// Создание узла с указателем на искомый узел
	newNode := CreateNode(data, node.Pointer) // &2
	// Установка указателя на новый узел
	(*node).Pointer = &newNode

	return linkedList
}

func InsertItemBeforeSelectNode(linkedList MyList, index int, data int) MyList {
	if index == 0 {
		return InsertItemFront(linkedList, data)
	}

	node := &(linkedList.Head)
	counter := 0

	// Итерация до искомого узла
	for counter < index-1 && node.Pointer != nil {
		node = node.Pointer // 3
		counter++
	}

	// Создание узла с указателем на следющий за искомым узел
	newNode := CreateNode(data, node.Pointer)
	// Установка указателя искомого узла на новый узел
	(*node).Pointer = &newNode

	return linkedList
}

func RemoveItem(linkedList MyList, index int) MyList {
	node := &(linkedList.Head)
	counter := 0

	if index == 0 {
		// Устанавливаем head на следующий узел
		return MyList{Head: *(node.Pointer)}
	}

	// Найти искомый узел
	for counter < index-1 && node.Pointer != nil {
		node = node.Pointer
		counter++
	}
	// У предыдущего узла установить указатель на последующий узел
	(*node).Pointer = node.Pointer.Pointer

	return linkedList
}
