package Array

import "fmt"

func Run() {
	array := CreateArray(10)
	fmt.Printf("%v\n", array)

	//// Добавление
	array = AppendItem(array, 1)
	fmt.Printf("%v\n", array)

	// Вставка
	array = InsertItem(array, 2, 0)
	fmt.Printf("%v\n", array)

	//// Удаление
	array = RemoveItemInArray(array, 0)
	fmt.Printf("%v\n", array)

	// Поиск по значению
	x := SelectItem(array, 1)
	fmt.Printf("%v\n", array[x])

	// Максимальное значение
	maxArray := Max(array)
	fmt.Printf("%v\n", maxArray)

	// Минимальное значение
	minArray := Min(array)
	fmt.Printf("%v\n", minArray)
}
