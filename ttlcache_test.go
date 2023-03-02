package fleckstoredb

import (
	"container/list"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntailizememory(t *testing.T) {
	cacheValue := NewTTlInit()

	Ok := assert.Equal(t, NewInit(), cacheValue)
	if !Ok {
		t.Errorf("The memoryInitialize failed")
	}
	if Ok {
		t.Log("The memoryInitialize passed")
	}
}
func TestSetcache(t *testing.T) {
	cachevalue := NewTTlInit()
	//This Level set the data in  front of the list
	Expeckdata := &Entry{
		Keys: []byte{118, 97, 108, 117, 101, 49},
	}
	checkvalue := cachevalue.setThecache(Expeckdata)
	assert.Equal(t, cachevalue.Doubly_list.Front(), checkvalue)
	//Again Reinitialize the Entry The error should
	//be thrown
	checkvalue = cachevalue.setThecache(Expeckdata)
	Ok := assert.Empty(t, checkvalue)
	if !Ok {
		t.Errorf("The Cache Check value failed")
	}
	if Ok {
		t.Error("The cache check Value passed")
	}

}

func TestBuffstate(t *testing.T) {
	cachevalue := NewTTlInit()
	//Now check Default limit Value is
	//Properly setted......
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 51, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 52, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 53, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 54, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 55, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 56, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 57, 32},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49, 48, 32},
		},
	}

	i := 0

	for {
		if i == 10 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}
	//Now add the extra Indexvalue into the queue
	//The list should Evict the item in the back
	value := new(list.Element)
	value = cachevalue.Doubly_list.Back()
	cachevalue.setThecache(&Entry{
		Keys: []byte{118, 97, 108, 117, 101, 49, 50, 32},
	})
	Ok := assert.NotEqualValues(t, value, cachevalue.Doubly_list.Back())
	if !Ok {
		t.Error("capacity test failed")

	}
	if Ok {
		t.Error("capacity  test passed")
	}
}
func TestGetcacheValueNil(t *testing.T, key []byte) {
	//Load The value into the queue
	cachevalue := NewTTlInit()

	i := 0
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
	}

	for {
		if i == 3 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}
	//If key is Not Present in the cache queue
	//It will throw the error
	Entrydata := cachevalue.getThecache([]byte{118, 97, 108, 117, 101, 49})
	Ok := assert.Nil(t, Entrydata)

	if !Ok {
		t.Error("The cache queue Check failed")
	}
	if Ok {
		t.Error("The cache queue check passed")
	}

}

func TestGetExactValue(t *testing.T) {
	cachevalue := NewTTlInit()
	Expect := []byte{118, 97, 108, 117, 101, 49}
	i := 0
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
	}

	for {
		if i == 3 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}
	returnEntry := cachevalue.getThecache(Expect)
	Ok := assert.Equal(t, Expect, returnEntry)
	if !Ok {
		t.Error("The Get access failed")
	}
	if Ok {
		t.Log("The Get access passed")
	}
}
func TestRemoveCacheNil(t *testing.T) {
	cachevalue := NewTTlInit()

	i := 0
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
	}

	for {
		if i == 3 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}
	//If key is Not Present in the cache queue
	//It will throw the error
	present := cachevalue.RemoveTheCache([]byte{118, 97, 108, 117, 101, 49})
	Ok := assert.Nil(t, present)

	if Ok {
		t.Log("The empty test passed")
	} else {
		t.Log("The empty check test failed")
	}
}
func TestRemoveExactvalue(t *testing.T) {
	cachevalue := NewTTlInit()
	ExpectValue := []byte{118, 97, 108, 117, 101, 49}
	i := 0
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
	}

	for {
		if i == 3 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}
	cachevalue.RemoveTheCache(ExpectValue)
	Datamatch := cachevalue.changeBufState(ExpectValue)
	FoundValue := new(list.Element)

	temp := cachevalue.Doubly_list.Front()
	for {
		if temp == nil {
			break
		}
		encodeCheck := cachevalue.changeBufState(temp.Value.(CacheHarmonize).EntryData.Keys)
		if encodeCheck == Datamatch {
			FoundValue = nil
			break
		}
		temp = temp.Next()
	}
	Ok := assert.Nil(t, FoundValue)
	if Ok {
		t.Log("The Remove Check failed")
	}
	if !Ok {
		t.Error("The Remove Check passed")
	}
}

func TestgetAllTheData(t *testing.T) {
	cachevalue := NewTTlInit()

	i := 0
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
	}

	for {
		if i == 3 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}
	valueCheck := cachevalue.GetAllTheData()
	Ok := assert.Equal(t, checkvalue, valueCheck)
	if Ok {
		t.Log("Test get all data Pass")
	}
	if !Ok {
		t.Error("Test get all data failed")
	}

}
func TestChangeBufstate(t *testing.T) {
	cacheValue := NewTTlInit()
	Expect := "Value1"
	value_Return := cacheValue.changeBufState([]byte(Expect))
	Ok := assert.Equal(t, Expect, value_Return)
	if Ok {
		t.Log("The changeBuftestPass")
	}
	if !Ok {
		t.Error("The Change log Test Failed")
	}
}
func TestClearstate(t *testing.T) {
	cachevalue := NewTTlInit()

	i := 0
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
	}

	for {
		if i == 3 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}
	cachevalue.ClearThecachedata()
	Ok := assert.Equal(t, list.New(), cachevalue.Doubly_list)
	if !Ok {
		t.Error("Clear cache Test failed")
	}
	if Ok {
		t.Log("Clear cache test passed")
	}
}
func TestSingleEvict(t *testing.T) {

	cachevalue := NewTTlInit()

	i := 0
	var checkvalue = []*Entry{
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 49},
		},
		&Entry{
			Keys: []byte{118, 97, 108, 117, 101, 50, 32},
		},
	}

	for {
		if i == 3 {
			break
		}
		cachevalue.setThecache(checkvalue[i])
		i++
	}

	Backvalue := new(list.Element)
	Backvalue = cachevalue.Doubly_list.Back()
	cachevalue.EvictTheSingledataOut()
	Ok := assert.NotEqual(t, Backvalue, cachevalue.Doubly_list.Back())
	if Ok {
		t.Log("The single Evict test passed")
	}
	if !Ok {
		t.Log("The Single Evict Test Failed")
	}

}

func TestEvictdataonExpiration(t *testing.T) {
	Default_time = 1
	var DataInit = InitializeTheEntry(&Entry{}, time.Duration(Default_time))
	time.Sleep(90 * time.Second)
	IsTrue := DataInit.IsJourneyEnded()
	var FrontValueInit = new(list.Element)
	if IsTrue {

		CacheValue := NewTTlInit()
		FrontValueInit = CacheValue.Doubly_list.PushFront(DataInit)
		CacheValue.EvictThedataOnExpiration(CacheValue.Doubly_list.Front())
		Ok := assert.NotEqual(t, FrontValueInit, CacheValue.Doubly_list.Front())
		if Ok {
			t.Log("The EvictData On expiration passed")
		}
		if !Ok {
			t.Log("The EvictData On Expiration failed")
		}

	}

}

func TestEvictDataOnRemoveDem(t *testing.T) {
	var CheckCacheValue = new(list.Element)
	var DataInit = InitializeTheEntry(&Entry{
		Keys: []byte{118, 97, 108, 117, 101, 49},
	}, time.Duration(Default_time))
	cachevalue := NewTTlInit()
	CheckCacheValue = cachevalue.Doubly_list.PushFront(DataInit)
	cachevalue.EvictTheOnRemoveDemand(cachevalue.Doubly_list.Front())
	Ok := assert.NotEqual(t, CheckCacheValue, cachevalue.Doubly_list.Front())
	if !Ok {
		t.Error("Test Evict On demand failed")
	}
	if Ok {
		t.Log("Test Evict On demand passed")
	}
}
