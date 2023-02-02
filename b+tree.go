package fleckstoredb

import "time"

type Bptree struct {
	Root      *Node
	TimeStamp time.Time
	FilePath  string
}

type Record struct {
	Value []byte
}

type Node struct {
	Keys    []byte
	Isleaf  bool
	Pointer []*interface{}
	Parent  *Node
	Next    *Node
	N       uint64
}
