package main

func main() {
	hastable := CreateHashTable()

	hastable.Insert(1, "a")
	hastable.Insert(1, "b")
	hastable.Del(2)
	hastable.Insert(1, "a")
	// key := 11
	// fmt.Printf("Value for key %d: %s", key, hastable.Get(key))
}
