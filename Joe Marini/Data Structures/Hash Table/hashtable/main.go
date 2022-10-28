package main

import (
	"fmt"
	"hash/fnv"
)

type Table struct {
	m     int
	table [][]kv
}

type kv struct {
	Key, Value string
}

func main() {
	t := New(1)
	t.Insert("foo", "bar")
	t.Insert("baz", "qux")
	t.Insert("foo", "escobar")
	//v, ok := t.Get("foo")
	//fmt.Printf("t.Get(%q) = (%q, %t)", "foo", v, ok)
	fmt.Println(t)

	fmt.Println(t.Get("foo"))
	fmt.Println(t.Get("baz"))
}

func New(m int) *Table {
	return &Table{
		m:     m,
		table: make([][]kv, m),
	}
}

func (t *Table) Get(key string) (string, bool) {
	i := t.hash(key)

	for j, kv := range t.table[i] {
		if key == kv.Key {
			// Found a match, return it!
			return t.table[i][j].Value, true
		}
	}

	return "", false
}

func (t *Table) Insert(key, value string) {
	i := t.hash(key)

	for j, kv := range t.table[i] {
		if key == kv.Key {
			// Overwrite previous value for the same key.
			t.table[i][j].Value = value
			return
		}
	}

	// Add a new value to the table.
	t.table[i] = append(t.table[i], kv{
		Key:   key,
		Value: value,
	})
}

func (t *Table) hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32()) % t.m
}
