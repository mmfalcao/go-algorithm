package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*Bucket
}

// bucket is a linked list in each slot of the hash table array
type Bucket struct {
	head *BucketNode
}

// bucketNode is a linked list node that holds the key
type BucketNode struct {
	key  string
	next *BucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insertBucket(key)
}

// Search will take ina key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].searchBucket(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].deleteBucket(key)
}

// InsertBucket will take in a key, create a node with the key and insert the node in the bucket
func (b *Bucket) insertBucket(k string) {
	if !b.searchBucket(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// Search will take in a key and return true if the bucket has
func (b *Bucket) searchBucket(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete will take in a key and delete the node from the bucket
func (b *Bucket) deleteBucket(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func main() {

	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("STAN")
	fmt.Println("STAN", hashTable.Search("STAN"))
	fmt.Println("KYLE", hashTable.Search("KYLE"))

	// fmt.Println(hash("RANDY"))

	testBucket := &Bucket{}
	testBucket.insertBucket("RANDY")
	testBucket.insertBucket("RANDY")
	testBucket.deleteBucket("RANDY")
	fmt.Println("RANDY", testBucket.searchBucket("RANDY"))
	fmt.Println("ERIC", testBucket.searchBucket("ERIC"))
}
