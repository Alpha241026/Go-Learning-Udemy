package main

import "fmt"

func main() {
	names := []string{"Alice", "John", "Mark"}
	fmt.Println(names)

	items := make([]int, 3, 5) //making integer slice having 3 slots with capacity for 5 (using make function)

	fmt.Printf("Items: +%v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)

	fmt.Printf("Items: +%v, Len: %d, Cap: %d\n", items, len(items), cap(items)) //adding elements beyond capacity will double size of slice
}
