package fleckstoredb

import "time"

const (
	//Key_is_deleted_on user_Request
	//It_is_not_deleted when cache is full
	//when_Cache_is_full automatically it will be deleted from the cache_data
	EvictReasonDeleted Evcitioneason = iota
	//Key_valid_time_is_expired
	//Then the eviction is setted to be Expired
	//Without_Expiration_the will not be deleted
	EvictReasonExpired
)

const (
	//LeastRecentaccessed scheme is setted
	//When The key is used recently by
	//user_accessed then It is setted to be LRA
	//Lra scheme is moved back in the cache List
	EvictionPolicyScheme EvictionPolicy = iota
)

type Evcitioneason int

func (s Evcitioneason) SetEvictionReason() string {

	return [...]string{0: "KeyDeleted", 1: "keyExpired"}[s]

}

type EvictionPolicy int

func (s EvictionPolicy) SetEvictionPolicy() string {
	return [...]string{0: "LRA"}[s]
}

//DoublyLinkedList Is used Here
//ratherThan the traditional singlyLinklist
//The cache willbe present On Each node
//We can Move Forward and backward the Cache Is present
//pointer is present for Next and PreviousNode
type Doubly_Node struct {
	Node                   *CacheHarmonize
	LeastRecentaccessedKey time.Duration
	Next                   *Doubly_Node
	Previous               *Doubly_Node
}
type Doubly_Link_List struct {
	//TheValue is pushed Into PriorityQueue
	//It is sorted in the order of low to high\
	//HEap sort is used for sorting
	CacheQueue *PriorityQueue
	Cache      map[int]*DoublyList
	headNode   *Doubly_Node
	TailNode   *DoublyList
}
