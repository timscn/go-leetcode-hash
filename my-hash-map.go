package main

type MyHashMap struct {
	m map[int]int
}

func ConstructorMyHashMap() MyHashMap {
	return MyHashMap{m: make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value

}

func (this *MyHashMap) Get(key int) int {
	if value, ok := this.m[key]; ok {
		return value
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.m, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
