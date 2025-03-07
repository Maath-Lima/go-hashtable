package main

import (
	"hashtable/utils"
)

var (
	defaultCapacity int16   = 5
	loadFactor      float32 = 0.7
	minLoadFactor   float32 = 0.25
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

func createHashTable(capacity int) *Hashtable {
	return &Hashtable{table: make([]Entry, capacity)}
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
	return (hashtable.table[index].key != 0 && !hashtable.table[index].tombstone)
}

func (hashtable *Hashtable) insert(key int, element string) {
	e, i := hashtable.lookUp(key)

	if e != nil && e.key == key {
		e.value = element
		e.tombstone = false
		return
	}

	hashtable.table[i] = Entry{key: key, value: element}
	hashtable.nElements++
}

func (hashtable *Hashtable) Insert(key int, element string) {
	hashtable.insert(key, element)
	hashtable.loadFactorResize()
}

func (hashtable *Hashtable) loadFactorResize() {
	var ht *Hashtable
	len := len(hashtable.table)

	balance := (float32(hashtable.nElements) / float32(len))

	// Do when delete method is up to use!
	if balance < minLoadFactor {
	}

	if balance > loadFactor {
		newLen := utils.FindNextPrime(len * 2)

		ht = createHashTable(newLen)

		for i := 0; i < len; i++ {
			if hashtable.isOccupied(i) {
				entry := hashtable.table[i]

				ht.insert(entry.key, entry.value)
			}
		}
	}

	if ht != nil && ht.nElements > 0 {
		*hashtable = *ht
	}
}

func (hashtable *Hashtable) Get(key int) string {

	e, _ := hashtable.lookUp(key)

	if e != nil {
		return e.value
	}

	panic("provided key doesn't exist")
}

func (hashtable *Hashtable) lookUp(key int) (*Entry, int) {
	capacity := len(hashtable.table)

	index := hashFunction(key, capacity)
	step := doubleHashFunction(key, capacity)

	iIndex := index
	var collisions int

	for hashtable.isOccupied(index) {
		if hashtable.table[index].key == key {
			return &hashtable.table[index], index
		}

		collisions++
		index = (iIndex + collisions*step) % capacity
	}

	return nil, index
}
