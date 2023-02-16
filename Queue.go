package fleckstoredb

import (
	"sync"
)

type Elements []CacheHarmonize

type PriorityQueue struct {
	// The Entry is added to the queue
	Entry *Elements
	Lock  *sync.Mutex
	// Multiple go routines are executed
	// One By One in different runtime executions
	wg *sync.WaitGroup
	//Ready For execution the value is 1
	//And not for execution and it will be setted to 0
	Execution chan bool
	//IsFleshOut ensures that all the elements are removed from the queue
	IsFleshOut bool
	//Flesh_In_And_Out_Lock Lock for Memory address of
	//Array of queue to do fleshout
	Flesh_In_Out_Lock sync.Mutex
}
// Intialize The Queue have to be prioritized 
func initialLizeThequeue(Size_Default int) *PriorityQueue {
	return &PriorityQueue{
		Entry:     make(CacheHarmonize,Size_Default), // Memory Leak Protection
		Lock:      &sync.Mutex{},
		Execution: make(chan bool, 1),
		Flesh_In_Out_Lock: &sync.Mutex{},
		wg: &sync.WaitGroup{},
	}
}
func (s *PriorityQueue) Enqueue(ch *CacheHarmonize)error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if len(*s.Entry) == nil{
         return nil
	}
    


}

//Heapify The array means move the sudden changes of an array value
//If if meets the certains logics
func (s *PriorityQueue) Heapify(Size, i int) {
	var left = 2*i + 1
	var right = 2*i + 2
	var largest = i
	if left < largest && s.Queue[left].End_Of_journey.Minute() > s.Queue[largest].End_Of_journey.Minute() {
		largest = left
	}
	if right < largest && s.Queue[right].End_Of_journey.Minute() > s.Queue[right].End_Of_journey.Minute() {
		largest = right
	}
	if largest != i {
		var temp = s.Queue[i]
		s.Queue[i] = s.Queue[largest]
		s.Queue[largest] = temp
		s.Heapify(Size, largest)
	}
}

// Get the Size of The array
func (s *PriorityQueue) GetTheValueofN() int {

	return len(s.Queue) / 2
}

// Tradtional HeapSort is used to sort
// Instead of using the quicksort or mergeSort

func (s *PriorityQueue) HeapSort() {
	N := s.GetTheValueofN()
	for i := N/2 - 1; i >= 0; i-- {
		s.Heapify(N, i)
	}
	for i := N - 1; i >= 0; i-- {
		var temp = s.Queue[i]
		s.Queue[i] = s.Queue[0]
		s.Queue[0] = temp
		s.Heapify(i, 0)
	}

}

// This function Swap the elements in the queue
func(s.)Swap(i, j int) {

}

//Pop removes the miniMalValue from the array of the data
//Means from the lowest to highest
//Pop is equal to remove
func (s *PriorityQueue) Pop() {

	s.InterChange(0)

}

// Interchange the value after the value is popped
// And intercahnge the value and store into the array
func (s *PriorityQueue) InterChange(i int) {
	var Ic []CacheHarmonize
	for j := i; j < len(s.Queue); j++ {
		if j == len(s.Queue)-1 {
			break
		}
		Ic[i] = s.Queue[j+1]
	}
	s.Queue = Ic

}
func (s *PriorityQueue) DisplayFront() *CacheHarmonize {

	return &s.Queue[0]

}
