package fleckstoredb

var (
	DefaultOrder = 10
)

type Btree struct {
	root     *Node
	filePath string
}
type Node struct {
	keys     []byte
	pointers []interface{}
	child    []*Node
	isLeaf   bool
	N        int
}
