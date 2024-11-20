package defines

// Pair represent key, value pair.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Pack create a pair with the given key and value
func Pack[K comparable, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{Key: key, Value: value}
}

// Unpack returns values contained in a pair.
func (t Pair[K, V]) Unpack() (K, V) {
	return t.Key, t.Value
}
