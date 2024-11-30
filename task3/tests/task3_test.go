package mymap

import (
	"fmt"
	"main/mymap"
	"strconv"
	"testing"
)

func TestBasicOperations(t *testing.T) {
	m := mymap.InitMyMap()

	// Add and Get
	m.Add("key1", 100)
	val, exists := m.Get("key1")
	if !exists || val != 100 {
		t.Errorf("expected key1 to have value 100, got %v, exists: %v", val, exists)
	}

	// Update and Get
	m.Add("key1", 200)
	val, exists = m.Get("key1")
	if !exists || val != 200 {
		t.Errorf("expected key1 to have value 200 after update, got %v, exists: %v", val, exists)
	}

	// Exists
	if !m.Exists("key1") {
		t.Error("expected key1 to exist")
	}
	if m.Exists("key2") {
		t.Error("did not expect key2 to exist")
	}

	// Remove and Get
	m.Remove("key1")
	val, exists = m.Get("key1")
	if exists {
		t.Errorf("expected key1 to be removed, but got value %v", val)
	}
}

func TestRehash(t *testing.T) {
	m := mymap.InitMyMap()

	// trigger rehash
	for i := 0; i < 500; i++ {
		m.Add(string("a"+strconv.Itoa(i)), i)
	}
	fmt.Println(m.Len())

	// Check elements
	for i := 0; i < 50; i++ {
		val, exists := m.Get(string("a" + strconv.Itoa(i)))
		if !exists || val != i {
			t.Errorf("expected key %s: value %d, but got %v: %v", string("a"+strconv.Itoa(i)), i, val, exists)
		}
	}
}

func TestCollisionHandling(t *testing.T) {
	m := mymap.InitMyMap()

	// Add keys with same hash
	m.Add("key1", 100)
	m.Add("keyA", 200)

	// Check if both are accessible
	val1, exists1 := m.Get("key1")
	val2, exists2 := m.Get("keyA")
	if !exists1 || val1 != 100 {
		t.Errorf("expected key1 to have value 100, got %v, exists: %v", val1, exists1)
	}
	if !exists2 || val2 != 200 {
		t.Errorf("expected keyA to have value 200, got %v, exists: %v", val2, exists2)
	}

	// Remove one and ensure the other still exists
	m.Remove("key1")
	if m.Exists("key1") {
		t.Error("expected key1 to be removed")
	}
	val, exists := m.Get("keyA")
	if !exists || val != 200 {
		t.Errorf("expected keyA to still exist with value 200, got %v, exists: %v", val, exists)
	}
}
