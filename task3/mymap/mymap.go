package mymap

import "crypto/sha256"

const mapStartBucketCount = 8

type MyMap struct {
	array     []*bucket
	countNode int
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func InitMyMap() MyMap {
	return MyMap{array: make([]*bucket, mapStartBucketCount), countNode: 0}
}

func (m *MyMap) Len() int {
	return len(m.array)
}

func (m *MyMap) needRehash() bool {
	if m.countNode == 0 {
		return false
	}
	load := m.countNode / len(m.array)
	return load > 6
}

func (m *MyMap) rehash() {
	newArray := make([]*bucket, len(m.array)*2)
	newMap := MyMap{array: newArray, countNode: 0}
	for i := 0; i < len(m.array); i++ {
		bucket := m.array[i]
		if bucket != nil && bucket.head != nil {
			nowNode := bucket.head
			newMap.Add(nowNode.key, nowNode.value)
			for nowNode.next != nil {
				nowNode = nowNode.next
				newMap.Add(nowNode.key, nowNode.value)
			}
		}
	}
	*m = newMap
}

func (m *MyMap) Add(key string, value int) {
	if m.needRehash() {
		m.rehash()
	}
	index := Hash(key, len(m.array))
	if m.array[index] == nil {
		m.array[index] = &bucket{}
	}
	if m.array[index].insert(key, value) {
		m.countNode++
	}
}

func (m *MyMap) Remove(key string) {
	index := Hash(key, len(m.array))
	if m.array[index] != nil {
		isDelete := m.array[index].delete(key)
		if isDelete {
			m.countNode--
		}
	}
}

func (m *MyMap) Copy() MyMap {
	copyArray := make([]*bucket, len(m.array))
	copy(copyArray, (m.array))
	newMap := MyMap{array: copyArray}
	return newMap
}

func (m *MyMap) Exists(key string) bool {
	index := Hash(key, len(m.array))
	if m.array[index] != nil {
		_, isHave := m.array[index].searchAndUpdate(key, nil)
		return isHave
	}
	return false
}

func (m *MyMap) Get(key string) (int, bool) {
	index := Hash(key, len(m.array))
	if m.array[index] != nil {
		return m.array[index].searchAndUpdate(key, nil)
	}
	return 0, false
}

func (b *bucket) insert(key string, value int) bool {
	_, ishave := b.searchAndUpdate(key, &value)
	if !ishave {
		node := bucketNode{key: key, value: value, next: b.head}
		b.head = &node
		return true
	}
	return false
}

func (b *bucket) searchAndUpdate(key string, value *int) (int, bool) {
	currentHead := b.head
	for currentHead != nil {
		if currentHead.key == key {
			if value != nil {
				currentHead.value = *value
			}
			return currentHead.value, true
		}
		currentHead = currentHead.next
	}
	return 0, false
}

func (b *bucket) delete(key string) bool {
	if b.head == nil {
		return false
	}
	if b.head.key == key {
		b.head = b.head.next
		return true
	}
	node := b.head
	for node.next != nil {
		if node.next.key == key {
			node.next = node.next.next
			return true
		}
		node = node.next
	}
	return false
}

func Hash(key string, size int) int {
	hash := sha256.New()
	hash.Write([]byte(key))
	byteHash := hash.Sum(nil)
	sum := 0
	for _, val := range byteHash {
		sum += int(val)
	}
	return sum % size
}
