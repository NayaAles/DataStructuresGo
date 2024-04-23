package SinglyLinkedList

import "fmt"

func Run() {
	linkedList := CrateLinkedList(5)
	PrintLinkedList(linkedList)

	// Добавление в начало
	linkedList = InsertItemFront(linkedList, 1)
	PrintLinkedList(linkedList)

	// Удаление узла
	linkedList = RemoveItem(linkedList, 5)
	PrintLinkedList(linkedList)

	// Поиск узла по значению
	node := SearchItem(linkedList, 1)
	fmt.Printf("Data: %d, Pointer: %p\n\n", node.Data, node.Pointer)

	// Добавление в конец
	linkedList = InsertItemTail(linkedList, 1)
	PrintLinkedList(linkedList)

	// Добавление после элемента по индексу
	linkedList = InsertItemAfterSelectNode(linkedList, 0, 2)
	PrintLinkedList(linkedList)

	// Добавление до элемента по индексу
	linkedList = InsertItemBeforeSelectNode(linkedList, 2, 3)
	PrintLinkedList(linkedList)
}
