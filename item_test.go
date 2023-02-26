package fleckstoredb

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	expect time.Duration
)

func TestEntryInitialize(t *testing.T) {

	value := InitializeTheEntry(&Entry{}, 1)
	expect = 1
	assert.Equal(t, expect, value.LifeTime)
	assert.Equal(t, &Entry{}, value.EntryData)

}
func TestTTlsetPeriod(t *testing.T) {
	expect = 3
	var dataInit = InitializeTheEntry(&Entry{}, expect)
	dataInit.sEtPeriod()
	CpuCount := runtime.NumCPU()
	if CpuCount > 1 {
		go assert.WithinDuration(t, time.Now().Add(3), dataInit.End_Of_journey, time.Minute)

	} else {
		assert.WithinDuration(t, time.Now().Add(3), dataInit.End_Of_journey, time.Minute)
	}
	//check for long haul means long journey
	dataInit.LifeTime = time.Hour
	dataInit.sEtPeriod()
	assert.WithinDuration(t, time.Now().Add(time.Hour), dataInit.End_Of_journey, time.Minute)
}
func TestJourneyEnd(t *testing.T) {
	// Check the TTl is expired
	expect = 2
	var dataInit = InitializeTheEntry(&Entry{}, expect)
	dataInit.sEtPeriod()
	time.Sleep(30)
	IS := dataInit.IsJourneyEnded()

	if IS {
		t.Log("The TImeLImits Is Expired")
	}
}
func TestTTlSet(t *testing.T) {
	var checkvalue bool

	expect = 2 * time.Minute
	var dataInit = InitializeTheEntry(&Entry{}, expect)
	time.Sleep(3 * time.Minute)
	// check whether Lifetimeis setted to 0
	// And is expired is marked to true
	dataInit.SetExpirationTime()
	checkvalue = assert.Equal(t, 0, dataInit.TimeTTlRemain)
	if checkvalue {
		checkvalue = assert.Equal(t, 1, dataInit.isExpired)
		if checkvalue {
			t.Log("TTL expiration set is valid")
		}
	} else {
		t.Log("TTL ExpitaionSet failed")
	}
	expect = 30 * time.Minute
	dataInit = InitializeTheEntry(&Entry{}, expect)
	dataInit.SetExpirationTime()
	checkvalue = assert.Equal(t, 35, dataInit.TimeTTlRemain)
	if checkvalue {
		checkvalue = assert.Equal(t, 0, dataInit.isExpired)
		if checkvalue {
			t.Log("TTL Expiration is valid in phase2 check")
		}
	} else {
	}
}
