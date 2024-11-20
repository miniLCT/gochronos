package common

import (
	"github.com/miniLCT/gochronos/defines"
)

// MapKeys 函数接收一个map[string]T类型的参数m，返回一个string类型的切片
func MapKeys[T any](m map[string]T) []string {
	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// MapValues 函数接收一个map类型的参数m，返回一个T类型的切片
func MapValues[T any](m map[string]T) []T {
	values := make([]T, 0)
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// MapMerge 合并两个map
func MapMerge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	m := make(map[K]V, len(m1)+len(m2))
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m
}

// Map2Pairs convert a map to a slice of pairs.
func Map2Pairs[K comparable, V any](m map[K]V) []defines.Pair[K, V] {
	pairs := make([]defines.Pair[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, defines.Pair[K, V]{Key: k, Value: v})
	}
	return pairs
}

// Pairs2Map convert a slice of pairs to a map.
func Pairs2Map[K comparable, V any](pairs []defines.Pair[K, V]) map[K]V {
	m := make(map[K]V, len(pairs))
	for _, p := range pairs {
		m[p.Key] = p.Value
	}
	return m
}
