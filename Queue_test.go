package fleckstoredb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	sampleEnrty    []*CacheHarmonize
	sampleExpected []*CacheHarmonize
)

func TestTheElementInit(t *testing.T) {

	values := NewInitPriorityQueue()
	assert.Equal(t, nil, values.Entry)
	assert.Equal(t, 1, values.Execution)
}

func TestPutELeMent(t *testing.T) {
	i := 0
	SampleCheck := time.Hour
	Values1 := NewInitPriorityQueue()
	for {
		if i == 10 {
			break
		}
		if i <= 3 {
			SampleCheck += 80
		}
		if i > 3 && i <= 5 {
			SampleCheck -= 30
		}
		if i > 5 && i <= 8 {
			SampleCheck += 30
		}
		Values := InitializeTheEntry(&Entry{}, SampleCheck)
		sampleEnrty = append(sampleEnrty, Values)

		go Values1.Entry.Enqueue(Values)
		i += 10

	}
	var Is = assert.Equal(t, sampleEnrty, Values1.Entry)
	if !Is {
		t.Error("The Stage1 is  SeTTed")
	}
	EntryValue := HeapSort()
	Is = assert.Equal(t, EntryValue, Values1.Entry)
	if Is {
		t.Log("The sorted queue is success")
	} else {
		t.Log("The Sorted Queue is Not Successfull")
	}
}
func HeapSort() []*CacheHarmonize {
	// The Value of N
	var N = len(sampleEnrty)
	N = N / 2

	for i := N; i <= 0; i-- {
		Heapify(N, i)
	}
	N = len(sampleEnrty)
	for j := N - 1; j <= 0; j-- {
		swap(0, j)
		Heapify(j, 0)
	}
	return sampleEnrty

}

func Heapify(N, i int) {
	var Larger = i
	var Left = 2*N + 1
	var Right = 2*N + 2

	if Larger < N && sampleEnrty[Left].LifeTime > sampleEnrty[Larger].LifeTime {
		Larger = Left

	}
	if Larger < N && sampleEnrty[Right].LifeTime > sampleEnrty[Larger].LifeTime {
		Larger = Right
		Heapify(N, Larger)
	}

	if Larger != i {
		swap(i, Larger)
	}
}

func swap(i, j int) {
	var sampleEnrtyS = sampleEnrty[i]
	sampleEnrty[i] = sampleEnrty[j]
	sampleEnrty[j] = sampleEnrtyS
}
func TestFirstElement(t *testing.T) {
	SampleCheck := time.Hour
	Values1 := NewInitPriorityQueue()
	var i = 0
	var Datas CacheHarmonize
	for {

		if i == 10 {
			break
		}
		if i <= 3 {
			SampleCheck += 80
		}
		if i > 3 && i <= 5 {
			SampleCheck -= 30
		}
		if i > 5 && i <= 8 {
			SampleCheck += 30
		}
		Values := InitializeTheEntry(&Entry{}, SampleCheck)
		if i == 1 {
			Datas = *Values
		}
		go Values1.Entry.Enqueue(Values)
		i += 10

	}

	assert.Equal(t, Values1.Entry.GetFront(), Datas)
}
func FlesHoUt(t *testing.T) {

	SampleCheck := time.Hour
	Values1 := NewInitPriorityQueue()
	var i = 0
	for {
		if i == 10 {
			break
		}
		if i <= 3 {
			SampleCheck += 80
		}
		if i > 3 && i <= 5 {
			SampleCheck -= 30
		}
		if i > 5 && i <= 8 {
			SampleCheck += 30
		}
		Values := InitializeTheEntry(&Entry{}, SampleCheck)

		go Values1.Entry.Enqueue(Values)
		i += 10

	}
	Values1.FLeshOut()
	var val = assert.Equal(t, nil, Values1)
	if val {
		t.Log("The Data is successfully flushedOut")
	}
}
func TestGetAllTheElement(t *testing.T) {
	i := 0
	SampleCheck := time.Hour
	Values1 := NewInitPriorityQueue()
	for {
		if i == 10 {
			break
		}
		if i <= 3 {
			SampleCheck += 80
		}
		if i > 3 && i <= 5 {
			SampleCheck -= 30
		}
		if i > 5 && i <= 8 {
			SampleCheck += 30
		}
		Values := InitializeTheEntry(&Entry{}, SampleCheck)
		sampleEnrty = append(sampleEnrty, Values)

		go Values1.Entry.Enqueue(Values)
		i += 10

	}
	var _, ele = Values1.GetallElements()
	var Check = assert.Equal(t, ele, sampleEnrty)
	if !Check {
		t.Log("The Elements are not equal inside the queue")
	}

}
