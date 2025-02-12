package main

var (
	defaultCapacity int16   = 5
	loadFactor      float32 = 0.7
)

type Entry struct {
	key       int
	value     string
	tombstone bool
}

type Hashtable struct {
	table     []Entry
	nElements int
}

func CreateHashTable() *Hashtable {
	return &Hashtable{table: make([]Entry, defaultCapacity)}
}

func hashFunction(key int, capacity int) int {
	return key % capacity
}

func doubleHashFunction(key int, capacity int) int {
	return 1 + (key % (capacity - 1))
}

func (hashtable *Hashtable) isOccupied(index int) bool {
	return (hashtable.table[index].key != 0 || hashtable.table[index].tombstone)
}

func (hashtable *Hashtable) Insert(key int, element string) {
	capacity := len(hashtable.table)

	index := hashFunction(key, capacity)
	step := doubleHashFunction(key, capacity)

	iIndex := index
	var collisions int

	for hashtable.isOccupied(index) {
		if hashtable.table[index].key == key {
			hashtable.table[index].value = element
			return
		}

		collisions++
		index = (iIndex + collisions*step) % capacity
	}

	hashtable.table[index] = Entry{key: key, value: element}
	hashtable.nElements++

	hashtable.loadFactorResize()
}

func (hashtable *Hashtable) loadFactorResize() {
	isLoadFactorTreshold := (hashtable.nElements / len(hashtable.table)) > int(loadFactor)

	if !isLoadFactorTreshold {
		return
	}

}
