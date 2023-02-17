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
	err, IS := dataInit.IsJourneyEnded()
	if err != nil {
		t.Errorf(err.Error())
	}
	if IS {
		t.Log("The TImeLImits Is Expired")
	}
}
