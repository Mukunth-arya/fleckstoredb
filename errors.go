package fleckstoredb

import "errors"

var (
	// Key is not found in the node
	ErrKeyNotFound = errors.New("the key not found")
	// The keys is required to be inserted in the node it should not be null
	ErrKeyRequired = errors.New("key required")
	//key is to larger than the given size
	ErrKeyIsLarge = errors.New("key is too large")
	//The Value is too large than the given size
	ErrValueisLarge = errors.New("the value is too large than the given size")
)
