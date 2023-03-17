package fleckstoredb

import "bytes"

const (

	//Default Order size
	DefaultOrderSize = 10
)

type EntryItem []Entry

//InsertAt the entry in the particular index
func (s *EntryItem) InsertAt(Index int, Entrydata *Entry) {
	var Empty Entry

	*s = append(*s, Empty)
	if Index < len(*s) {
		copy((*s)[Index+1:], (*s)[Index:])

	}
	(*s)[Index] = *Entrydata
}

//Remove the entry at the particular Index
func (s *EntryItem) RemoveAt(Index int) {
	if Index < len(*s) {
		copy((*s)[Index:], (*s)[Index+1:])
	}
	//Restore the final index to be  empty
	var Empty Entry
	(*s)[len(*s)-1] = Empty

}

//Truncate means remove the certain portion of data
//Front and Back
func (s *EntryItem) Truncate(Index int) {
	var Datas []Entry
	if Index < len(*s) {
		copy((*s), (*s)[:Index])
	}
	Datas = (*s)[Index+1:]
	var Empty Entry
	i := 0
	for {
		if i == len(Datas) {
			break
		}
		Datas[i] = Empty
		i++
	}
}

//Find The data In the list Of EntryItems
//Entrydata and the indexvalue
func (s *EntryItem) FinDEntry(EntryData *Entry) (IndexValue int, Isfound bool) {
	var i = 0
	for {
		if i == len((*s)) {
			break
		}
		BuFfValue1 := BuFfConverter((*s)[i].Keys)
		BuFfValue2 := BuFfConverter(EntryData.Keys)
		if BuFfValue1 == BuFfValue2 {
			return i, true
		} else {
			return -1, false
		}

	}

}

//Periodically the last item will delist from the list of item
//and last item will be freed
func (s *EntryItem) Pop() *Entry {
	Popvalue, PopEntryItem := len(*s)-1, (*s)[len(*s)-1]
	//Set The last Value To Empty
	var Empty Entry
	(*s)[Popvalue] = Empty
	return &PopEntryItem
}

//Buffer converter Convert  the stream of bytes
func BuFfConverter(ConvertValue []byte) string {

	BuFfValue1 := bytes.NewBuffer(ConvertValue).String()

	return BuFfValue1
}

//Node is The single Entry in  entire list
type Node struct {
	//Each node has It's own Entry_Item
	//And It has pointer to array of Children Nodes
	Key_Value_store EntryItem

	Children []*Node
	// is_leaf ensures That the Node has the children Node
	isLeaf bool
	//Each Node has It's own memory address associated
	//Program counter ensures each time enable triggered
	//from the volatile-memory collective data have it's own memory address
	Address int64
}
