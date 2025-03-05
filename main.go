package main

import (
	"fmt"
)

func main() {
	hastable := CreateHashTable()

	hastable.Insert(1, "a")
	hastable.Insert(6, "b")
	hastable.Insert(11, "c")
	hastable.Insert(16, "d")

	key := 11
	fmt.Printf("Value for key %d: %s", key, hastable.Get(key))
}
