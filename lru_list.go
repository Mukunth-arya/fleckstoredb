package fleckstoredb

type NodeElement struct {

	// Doubly LinkList is used to implement the cache
	// Rather Than the traditional LinkList
	// The Doubly LinkList Holds the Next and previous pointers
	Next, Prev *NodeElement
	//The Value Could Be any type of datastructure
	Value      any
	DoublyList *DoublyList
}
type DoublyList struct {
	// Each Time New Node is Initalized the length will be incremented
	len uint64
	//Root of  the element
	Root NodeElement
}
