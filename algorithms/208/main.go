package main

import "fmt"

// Trie Tree
type Trie struct {
	IsFinish bool
	Child    []*Trie
}

// Constructor 新建一个 Root 节点
func Constructor() Trie {
	return Trie{IsFinish: false, Child: make([]*Trie, 26)}
}

// Insert 增加一个字符串
func (t *Trie) Insert(word string) {
	cur := t
	for i, v := range word {
		if cur.Child[v-'a'] == nil {
			child := Trie{
				IsFinish: i == len(word),
				Child:    make([]*Trie, 26),
			}
			cur.Child[v-'a'] = &child
			cur = &child
		} else {
			cur = cur.Child[v-'a']
		}
	}
	cur.IsFinish = true
}

// Search 完整匹配字符串
func (t *Trie) Search(word string) bool {
	cur := t
	for _, v := range word {
		if cur.Child[v-'a'] != nil {
			cur = cur.Child[v-'a']
		} else {
			return false
		}
	}
	return cur.IsFinish
}

// StartsWith 搜索起始字符串
func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for _, v := range prefix {
		if cur.Child[v-'a'] != nil {
			cur = cur.Child[v-'a']
		} else {
			return false
		}
	}
	return true
}

func main() {
	obj := Constructor()
	obj.Insert("word")
	obj.Insert("prefix")
	fmt.Println(obj)

	fmt.Println(obj.Search("word"), obj.StartsWith("pre"))
}
