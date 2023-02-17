package fleckstoredb

import (
	"errors"
	"time"
)

var (
	ErrPeriodSet = errors.New("TTlSetLocked")
)

type CacheHarmonize struct {
	EntryData      *Entry
	LifeTime       time.Duration
	End_Of_journey time.Time
	Stuck_here     chan bool
}

func InitializeTheEntry(entry *Entry, TTl time.Duration) *CacheHarmonize {
	value := &CacheHarmonize{
		EntryData:  entry,
		LifeTime:   TTl,
		Stuck_here: make(chan bool, 1),
	}
	go value.sEtPeriod()
	return value
}
func (s *CacheHarmonize) sEtPeriod() error {
	if s.LifeTime/1 == 0 && s.LifeTime < 0 {
		return ErrPeriodSet
	}
	s.End_Of_journey = time.Now().Add(s.LifeTime)
	return nil
}
func (s *CacheHarmonize) IsJourneyEnded() (error, bool) {
	if s.LifeTime/1 == 0 && s.LifeTime < 0 {
		return ErrPeriodSet, false
	}
	return nil, s.End_Of_journey.Before(time.Now())

}
