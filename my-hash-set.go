package main

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

type MyHashSet struct {
	data []int
}

func Constructor() MyHashSet {
	return MyHashSet{
		data: []int{},
	}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.data = append(this.data, key)
	}

}

func (this *MyHashSet) Remove(key int) {
	for index, value := range this.data {
		if value == key {
			this.data = append(this.data[:index], this.data[index+1:]...)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, e := range this.data {
		if e == key {
			return true
		}
	}
	return false
}
