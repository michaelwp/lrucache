package lrucache

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	lru := NewLRUCache(1)
	lru.Put("one", 10)

	tests := []struct {
		name     string
		lruCache LRUCache
		queryKey string
		expect   any
	}{
		{
			name:     "data available",
			lruCache: lru,
			queryKey: "one",
			expect:   10,
		},
		{
			name:     "data unavailable",
			lruCache: lru,
			queryKey: "two",
			expect:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.lruCache.Get(tt.queryKey)
			if actual != tt.expect {
				t.Errorf("expect: %v, actual:%v", tt.expect, actual)
			}
		})

		return
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value any
	}{
		{
			name:  "successfully put new data",
			key:   "one",
			value: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := NewLRUCache(1)
			lru.Put(tt.key, tt.value)

			value := lru.Get(tt.key)
			if value != tt.value {
				t.Errorf("expect: %v, actual:%v", tt.value, value)
			}
		})

		return
	}
}

func TestViewList(t *testing.T) {
	type data struct {
		key   string
		value any
	}

	tests := []struct {
		name     string
		capacity int
		data     []*data
		expect   []*data
	}{
		{
			name:     "successfully view all list of data",
			capacity: 3,
			data:     []*data{{"one", 10}, {"two", 20}, {"three", 30}},
			expect:   []*data{{"three", 30}, {"two", 20}, {"one", 10}},
		},
		{
			name:     "successfully view list of data, with some data are nil",
			capacity: 5,
			data:     []*data{{"one", 10}, {"two", 20}, {"three", 30}},
			expect:   []*data{{"three", 30}, {"two", 20}, {"one", 10}, nil, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := NewLRUCache(tt.capacity)
			for _, d := range tt.data {
				lru.Put(d.key, d.value)
			}

			actual := lru.ViewList()
			for i, d := range actual {
				if d == nil {
					continue
				}

				expectKeyValue := fmt.Sprintf("%s:%v", tt.expect[i].key, tt.expect[i].value)
				actualKeyValue := fmt.Sprintf("%s:%v", d.Key, d.Value)

				if actualKeyValue != expectKeyValue {
					t.Errorf("expect: %v, actual:%v", expectKeyValue, actualKeyValue)
				}
			}
		})

		return
	}
}
