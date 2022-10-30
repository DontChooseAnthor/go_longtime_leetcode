package main

import "fmt"

// ***时间复杂度***

// 数组按照下标查找为O(1)
// 数组顺序查找为O(n)
// 数组二分查找为O(logn)
// 数组删除插入为O(n)

// 线性链表处理节点为O(1)
// 线性链表遍历查找为O(n)

// 哈希函数，元素通过该函数得到存储地址（哈希运算）
// 哈希冲突（碰撞），数组长度有限的情况下，存储地址出现冲突

// 主干Entry函数，主干长度一定是2的次幂，主干存的是k-v
type KV struct {
	Key   string
	Value string
}

// k-v存在链表中
type LinkNode struct {
	Data     KV
	NextNode *LinkNode
}

func CreateLink() *LinkNode {
	return &LinkNode{KV{
		Key:   "",
		Value: "",
	}, nil}

}

// 尾插法
func (link *LinkNode) AddNode(data KV) int {
	var count = 0
	tail := link
	for {
		count += 1
		if tail.NextNode == nil {
			break
		} else {
			tail = tail.NextNode
		}
	}
	var newNode = &LinkNode{data, nil}
	tail.NextNode = newNode
	return count + 1
}

const BucketCount = 16

type HashMap struct {
	Buckets [BucketCount]*LinkNode
}

func CreateHashMap() *HashMap {
	myMap := &HashMap{}
	for i := 0; i < BucketCount; i++ {
		myMap.Buckets[i] = CreateLink()
	}

	return myMap
}

func HashCode(key string) int {
	var sum = 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return (sum % BucketCount)
}

func (myMap *HashMap) AddKeyValue(key string, value string) {
	var mapIndex = HashCode(key)

	var link = myMap.Buckets[mapIndex]
	if link.Data.Key == "" && link.NextNode == nil {
		link.Data.Key = key
		link.Data.Value = value
		fmt.Printf("node key:%v add to buckets %d first node \n", key, mapIndex)
	} else {
		index := link.AddNode(KV{key, value})
		fmt.Printf("node key:%v add to buckets %d %dth node \n", key, mapIndex, index)
	}
}

func (myMap *HashMap) GetValueForKey(key string) string {
	var mapIndex = HashCode(key)
	var link = myMap.Buckets[mapIndex]
	var value string

	head := link
	for {
		if head.Data.Key == key {
			value = head.Data.Value
			break
		} else {
			head = head.NextNode
		}
	}
	return value
}

func main() {
	myMap := CreateHashMap()
	myMap.AddKeyValue("001", "1")
	myMap.AddKeyValue("002", "2")
	myMap.AddKeyValue("003", "3")
	myMap.AddKeyValue("004", "4")
	myMap.AddKeyValue("005", "5")
	myMap.AddKeyValue("006", "6")
	myMap.AddKeyValue("007", "7")
	myMap.AddKeyValue("008", "8")
	myMap.AddKeyValue("009", "9")
	myMap.AddKeyValue("010", "10")
	myMap.AddKeyValue("011", "11")
	myMap.AddKeyValue("012", "12")
	myMap.AddKeyValue("013", "13")
	myMap.AddKeyValue("012", "14")
	myMap.AddKeyValue("015", "15")
}
