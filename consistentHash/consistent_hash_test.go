package consistentHash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewConsistentHash(t *testing.T) {
	hash := NewConsistentHash(3, nil)
	hash.Add("localhost:9999", "localhost:9998", "localhost:9997")

	t.Log(hash.Get("A"))
	t.Log(hash.Get("B"))
	t.Log(hash.Get("C"))
	t.Log(hash.Get("D"))
	t.Log(hash.Get("E"))
	t.Log(hash.Get("F"))
}

func TestHashing(t *testing.T) {
	hash := NewConsistentHash(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	fmt.Println(">>>>>>>>>>>>>>>>>>>")
	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}
	fmt.Println(hash.hashMap)
	fmt.Println(hash.keys)
	for k, v := range testCases {
		fmt.Println(hash.Get(k), v)
		if hash.Get(k) != v {
			t.Errorf("1 Asking for %s, should have yielded %s", k, v)
		}
	}

	// Adds 8, 18, 28
	hash.Add("8")
	hash.Del("8")
	fmt.Println(hash.hashMap)
	fmt.Println(hash.keys)
	// 27 should now map to 8.
	//testCases["27"] = "8"
	//for k, v := range testCases {
	//	fmt.Println(hash.Get(k), v)
	//	if hash.Get(k) != v {
	//		t.Errorf("Asking for %s, should have yielded %s", k, v)
	//	}
	//}
	for k, v := range testCases {
		fmt.Println(hash.Get(k), v)
		if hash.Get(k) != v {
			t.Errorf("2 Asking for %s, should have yielded %s", k, v)
		}
	}
}
