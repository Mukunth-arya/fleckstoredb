package fleckstoredb

import (
	"errors"
	"sync"
)

type Elements []CacheHarmonize

var (
	ErrProcessQueue  = errors.New("Unable to Process Queue")
	ErrLoadElements  = errors.New("Unable to load the elements inside the queue")
	ErrEmptyElements = errors.New("Unable to retrive the elements")
)

type PriorityQueue struct {
	// The Entry is added to the queue
	Entry *Elements
	Lock  *sync.Mutex

	//Ready For execution the value is 1
	//And not for execution and it will be setted to 0
	Execution chan bool
	//IsFleshOut ensures that all the elements are removed from the queue
	IsFleshOut bool
	//Flesh_In_And_Out_Lock Lock for Memory address of
	//Array of queue to do fleshout
	Flesh_In_Out_Lock sync.Mutex
}

//Load the data into The queue
//And sort the data in the queue
//Here traditional heap sort is used here
func (s *Elements) Enqueue(Element *CacheHarmonize) {
	var (
		entry int
	)

	*s = append(*s, *Element)
	// Length of the given element
	entry = s.Len()
	entry = entry / 2
	if entry >= 1 {

		for i := entry; i >= 0; i-- {
			s.Heapify(entry, i)
		}
		entry = s.Len()
		for i := entry - 1; i >= 0; i-- {
			//After HEapify The Larger data is Moved to begining
			//of the index and Now swap it to End of the array
			s.Swap(0, i)
			//Heapify Again Rest of the elements is swapped
			s.Heapify(i, 0)
		}

	}

}

//Heap sort is used for cache efficient Model rather than the
//traditional quicksort or mergesort is not efficient for cache management
//HeapSort has an time complexity of 0(NlogN)
func (s *Elements) Heapify(n, i int) {
	var Larger = i
	var Left = 2*n + 1
	var Right = 2*n + 2

	if Larger < n && (*s)[Left].LifeTime > (*s)[i].LifeTime {
		Larger = Left
	}
	if Larger < n && (*s)[Right].LifeTime > (*s)[i].LifeTime {
		Larger = Right
	}
	//compare the data that the larger
	if Larger != i {
		s.Swap(i, Larger)
		//Go it recursively for muchMore sort
		s.Heapify(n, Larger)
	}

}

//The length of the Givenqueue
func (s *Elements) Len() int {
	return len(*s)
}

// Swap The Elements Inside The queue
//If It meets Required Conditions
func (s *Elements) Swap(i, j int) {
	var eleMent = (*s)[i]
	(*s)[i] = (*s)[j]
	(*s)[j] = eleMent
}

//Get all  the elements from the Queue
//And Empty The queue from By Popping the elements
func (s *Elements) GetAll() (error, []*CacheHarmonize) {

	len := s.Len()
	var GetAll = make([]*CacheHarmonize, len)
	if len == 0 {
		return ErrProcessQueue, nil
	}
	for i := 0; i <= len-1; i++ {
		GetAll = append(GetAll, s.GetFront())
	}
	return nil, GetAll
}

//Pop the elements Means Getting the elements one by one
//From the sorted array of queue
//Means popping in the frontWords
func (s *Elements) Pop() {

	var len = s.Len()
	var Reininitalize = make([]CacheHarmonize, len)
	if len != 0 {

	}
	for i := 0; i < len; i++ {
		if i == len-1 {
			break
		}
		Reininitalize = append(Reininitalize, (*s)[i+1])
	}
	*s = Reininitalize
	return
}

//Get the elements from the first Index
func (s Elements) GetFront() *CacheHarmonize {
	return &(s)[0]
}

// Initialize the newPriority queue
func NewInitPriorityQueue() *PriorityQueue {
	ValueInit := &PriorityQueue{
		Entry:     nil,
		Execution: make(chan bool, 1),
	}
	return ValueInit
}

//Put the array of elements in the queue
//And The Elements is sorted based on the TTL
func (s *PriorityQueue) Put(Entry ...CacheHarmonize) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if (*s).IsFleshOut == false {
		return ErrLoadElements
	}
	if s.Entry == nil {

		for _, i := range Entry {

			s.Entry.Enqueue(&i)
		}

	}
	return nil
}

//TO enqueue the all the elements from the queue
//Means setting the state of the queue in the zerostate
func (s *PriorityQueue) FLeshOut() {
	s.Flesh_In_Out_Lock.Lock()
	defer s.Flesh_In_Out_Lock.Unlock()
	if (*s).IsFleshOut == true {
		(*s).Entry = nil
		(*s).IsFleshOut = true

	}

}

// Get the empty all the elements from the queue
// And the set the queue in the empty state
func (s *PriorityQueue) GetallElements() (error, []*CacheHarmonize) {
	s.Lock.Lock()
	s.Lock.Unlock()
	err, ValueELements := s.Entry.GetAll()
	if err != nil {
		return err, nil
	}
	return nil, ValueELements
}

//Get the elements First iNdex elementS from the queue
//Means the highest of the element
func (s *PriorityQueue) GetFrontElement() *CacheHarmonize {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	return (*s).Entry.GetFront()
}
