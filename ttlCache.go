package fleckstoredb

import (
	"bytes"
	"container/list"
	"errors"
	"runtime"
	"sync"
	"time"
)

var (
	// one hour five minutes is the default time remian for the data to be
	// in the ttl_cache if it exceeds with the given condition and it will be evicted with certain reason
	Default_time time.Duration // 3.9e+6 in seconds
)

const (

	//Key_is_deleted_on user_Request
	//when_Cache_is_full automatically it will be deleted from the cache_data
	EvictReasonDeleted Evcitioneason = iota
	//when Key_valid_time_is_expired
	//Then the eviction is setted to be Expired
	//Without an Expiration_the will not be deleted
	EvictReasonExpired
	//Average size for an cache is 20 kb
	//So Only 10 element is fixed in cache queue
	//This value will not incremented
	Default_size = 10
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

	return [...]string{1: "keydeleted", 2: "keyexpired"}[s]

}

type EvictionPolicy int

func (s EvictionPolicy) SetEvictionPolicy() string {
	return [...]string{1: "LRA"}[s]
}

//TTl_cache elements are inserted and deleted
//On user request...
type TTL_Cache struct {
	mu           sync.Mutex
	List_element map[string]*Entry
	//Doubly_link_list is used here rather than
	//the singly linklist the cache is present
	//from head to tail node each node has pointer to next Node
	Doubly_list *list.List
	//The number of element is  inserted into the cache
	//Each time the new element is  inserted The the value is  incremented
	N        int
	eviction *Eviction
}
type Eviction struct {
	mu sync.Mutex
	// During Eviction the record of the data will be maintained and verified
	//And the Eviction reason will be update
	EvictionMap map[string]any
}

//Initialize the memory For Large data entry to the queue
func NewTTlInit() *TTL_Cache {

	cache := &TTL_Cache{
		List_element: make(map[string]*Entry, 10),
		Doubly_list:  list.New(),

		eviction: &Eviction{
			EvictionMap: make(map[string]any, 10),
		},
	}

	return cache
}

func (s *TTL_Cache) setThecache(TTlEntry *Entry) *CacheHarmonize {

	// check the data is present in the queue if it's present then the entry will be not
	// permitted....
	BigEndian := s.changeBufState(TTlEntry.Keys)
	s.mu.Lock()
	Present := s.List_element[BigEndian]
	s.mu.Unlock()
	// if key is present the data Will not be setted in the cache
	if Present != nil {
		return nil
	}
	s.List_element[BigEndian] = TTlEntry
	//EachTimeTheNewEntry comes the queue will be cleared from the back
	if s.N == Default_size {
		s.EvictTheSingledataOut()
	}
	//Initialize The Cache-Harmonize Value
	Default_time = 65
	value := InitializeTheEntry(TTlEntry, Default_time)

	s.mu.Lock()
	s.Doubly_list.PushFront(value)
	//Single_data is inserted So increment the value
	s.N++
	s.mu.Unlock()
	return value
}

//Get the single_data out from the cache_value
//setted as an lra and pushed back
func (s *TTL_Cache) getThecache(key []byte) *Entry {

	//Get the Encoded the into uint64
	Decodedvalue := s.changeBufState(key)
	//Check The data is Not Present
	//Then through it out  a error
	s.mu.Lock()
	Present := s.List_element[Decodedvalue]
	s.mu.Unlock()
	if Present == nil {
		errors.New("Key is not present")
		return nil
	}
	temp := s.Doubly_list.Front()
	for {
		if temp == nil {
			break
		}
		encodeCheck := s.changeBufState(temp.Value.(CacheHarmonize).EntryData.Keys)
		if encodeCheck == Decodedvalue {
			Cd := temp.Value.(CacheHarmonize)
			if !Cd.IsJourneyEnded() {

				go func() {
					s.mu.Lock()
					swap1 := temp
					s.Doubly_list.MoveToBack(swap1)
					temp = nil

					s.mu.Unlock()
				}()
				return Present
				break
			}
			if Cd.IsJourneyEnded() {
				s.EvictThedataOnExpiration(temp)
			}

		}
		temp = temp.Next()
	}
	return nil
}

//Remove The data from the queue on user demand
//And set the It will preceed with the eviction policy

func (s *TTL_Cache) RemoveTheCache(key []byte) *Entry {
	// check the data is present in the queue if it's present then the entry will be not
	// permitted....
	BigEndian := s.changeBufState(key)
	s.mu.Lock()
	Present := s.List_element[BigEndian]
	s.mu.Unlock()
	// if key is present the data Will not be setted in the cache
	if Present == nil {
		return nil
	}
	temp := s.Doubly_list.Front()
	for {
		if temp == nil {
			break
		}
		encodeCheck := s.changeBufState(temp.Value.(CacheHarmonize).EntryData.Keys)
		if encodeCheck == BigEndian {
			s.EvictTheOnRemoveDemand(temp)
			break
		}
		temp = temp.Next()
	}
	return Present
}

// Get All The Element out of cache queue
//Reset The entireScheme and entire Value
func (s *TTL_Cache) GetAllTheData() (datas []CacheHarmonize) {
	temp := s.Doubly_list.Front()

	for {
		if temp == nil {
			break
		}
		datas = append(datas, temp.Value.(CacheHarmonize))
		temp = temp.Next()
	}
	//Reset the entire list
	s.Doubly_list = list.New()
	return

}

func (s *TTL_Cache) changeBufState(key []byte) string {

	value := bytes.NewBuffer(key).String()
	return value
}

//EvictThedataFromtheQueue when the capacity is reached and
//element from the back will be removed rather than the front
func (s *TTL_Cache) EvictTheSingledataOut() {
	var wg sync.WaitGroup
	//Remove the data from the back
	//Because It is matched with LRA policy
	s.mu.Lock()

	s.eviction.EvictionMap[EvictReasonDeleted.SetEvictionReason()] = s.Doubly_list.Back()
	s.mu.Unlock()

	// It will be executed Only-when
	// the multithreaded execution means more number of cpu
	if runtime.NumCPU() > 1 {
		go func() {
			s.mu.Lock()
			wg.Add(1)
			s.Doubly_list.Remove(s.Doubly_list.Back())
			s.N--
			wg.Done()
			s.mu.Unlock()
		}()
		wg.Wait()
	}
}

//Evict the data on Expiration
func (s *TTL_Cache) EvictThedataOnExpiration(Values *list.Element) {
	var wg sync.WaitGroup
	//Evict The single data out of the queue on expiration...
	s.mu.Lock()
	s.eviction.EvictionMap[EvictReasonExpired.SetEvictionReason()] = Values
	s.mu.Unlock()
	// It will be executed Only-when
	// the multithreaded execution means more number of cpu
	if runtime.NumCPU() > 1 {
		go func() {
			s.mu.Lock()
			wg.Add(1)

			s.List_element[string(Values.Value.(CacheHarmonize).EntryData.Keys)] = nil
			s.Doubly_list.Remove(Values)
			s.N--
			wg.Done()
			s.mu.Unlock()
		}()
		wg.Wait()
	}
	s.List_element[string(Values.Value.(CacheHarmonize).EntryData.Keys)] = nil
	s.Doubly_list.Remove(Values)

}

//Evict The data On the user demand
//Eviction Policy is preceeded
func (s *TTL_Cache) EvictTheOnRemoveDemand(elem *list.Element) {

	var wg sync.WaitGroup
	//Evict The single data out of the queue on expiration...
	s.mu.Lock()
	s.eviction.EvictionMap[EvictReasonDeleted.SetEvictionReason()] = elem
	s.mu.Unlock()
	// It will be executed Only-when
	// the multithreaded execution means more number of cpu
	if runtime.NumCPU() > 1 {
		go func() {
			s.mu.Lock()
			wg.Add(1)
			s.List_element[string(elem.Value.(CacheHarmonize).EntryData.Keys)] = nil
			s.Doubly_list.Remove(elem)
			s.N--
			wg.Done()
			s.mu.Unlock()
		}()
		wg.Wait()
	}
	s.Doubly_list.Remove(elem)
	s.List_element[string(elem.Value.(CacheHarmonize).EntryData.Keys)] = nil

}

//If_admin_command send the all the element want to remove from the
//cache or clear the entire state this method will pop up
func (s *TTL_Cache) ClearThecachedata() {
	s.mu.Lock()
	s.Doubly_list = list.New()
	//Change The Number oF entry to nil state
	s.List_element = make(map[string]*Entry)
	s.N = 0
	s.mu.Unlock()

}
