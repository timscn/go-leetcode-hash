package main

import (
	"fmt"
	"reflect"
	"sort"
)

func emptyMapPrint() {
	hashTable := make(map[string]int)
	fmt.Println("hashTable is type of: ", reflect.TypeOf(hashTable))
	fmt.Printf("Empty M: %v\n", hashTable)
	hashTable["route"] = 66
	hashTable["road"] = 2342
	fmt.Printf("non-empty M: %v\n", hashTable)
	fmt.Printf("value for dsf is :%v in our HashTable: \n", hashTable["dsf"])
	delete(hashTable, "dsf")
	for key, value := range hashTable {
		fmt.Printf("The key: %v and the value: %v\n", key, value)
	}

	// Map is not guarantee return order so
	keys := []string{}
	for key, _ := range hashTable {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for key, v := range keys {
		fmt.Println("Key:", key, "Value:", hashTable[v])
	}
}
