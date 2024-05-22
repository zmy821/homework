package main

import (
	"fmt"
)

func DeleteElement[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]

	newSlice := make([]T, len(slice))
	copy(newSlice, slice)

	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	indexDelete := 2
	fmt.Printf("原始长度: %d,原始容量: %d\n", len(slice), cap(slice))
	fmt.Println("初始切片:", slice)
	newSlice := DeleteElement(slice, indexDelete)
	fmt.Println("删除特定元素后的切片:", newSlice)
	fmt.Printf("删除后长度:%d ,删除后容量: %d\n", len(newSlice), cap(newSlice))

}
