package common

import (
	"sort"
)

// SliceContains 从切片中判断是否存在
func SliceContains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// SliceRmDup 删除切片中的重复项
func SliceRmDup[T comparable](slice []T) []T {
	m := map[T]struct{}{}
	var result []T
	for _, s := range slice {
		if _, ok := m[s]; !ok {
			result = append(result, s)
			m[s] = struct{}{}
		}
	}
	return result
}

func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// SortString sorts []string in increasing order
func SortString(a []string) { sort.Slice(a, func(i, j int) bool { return a[i] < a[j] }) }

// ToSet returns a single slice containing the unique values from one or more slices.
// The order of the items in the result is not guaranteed.
func ToSet[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	m := map[T]struct{}{}
	for _, slice := range slices {
		for _, value := range slice {
			m[value] = struct{}{}
		}
	}
	result := []T{}
	for k := range m {
		result = append(result, k)
	}
	return result
}

func SliceIsSubset[T comparable](small, big []T) bool {
	m := make(map[T]bool)
	for _, val := range big {
		m[val] = true
	}
	for _, val := range small {
		if !m[val] {
			return false
		}
	}

	return true
}

// Reverse means the first becomes the last, the second becomes the second to last, and so on
func Reverse[T any](s []T) []T {
	l := len(s)
	mid := l >> 1

	for i := 0; i < mid; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	return s
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// DifferenceSet returns the elements in `a` that aren't in `b`.
func DifferenceSet[T comparable](a, b []T) []T {
	mp := make(map[T]struct{}, len(b))
	for _, v := range b {
		mp[v] = struct{}{}
	}

	res := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := mp[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
