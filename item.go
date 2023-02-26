package fleckstoredb

import (
	"errors"
	"time"
)

var (
	ErrPeriodSet = errors.New("TTlSetLocked")
)

const (
	//one hour 5 minutes is an average ttl lifetime
	TTL_lifeTime = 3.9e+6
)

type CacheHarmonize struct {
	EntryData *Entry
	// LifeTime is meant to be TTL
	//Amount of time that the data is lived in volatile state
	LifeTime time.Duration
	//End_of_journey is meant to be
	//current_time + TTl time
	End_Of_journey time.Time
	//Stuck_here means Data to be settled
	//for entire_lifetime means 1 hour
	Stuck_here chan bool
	//Is_expired_is_setted when the data's lifetime is expired
	//Is compared with the current_time
	isExpired chan bool

	TimeTTlRemain time.Duration
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
func (s *CacheHarmonize) IsJourneyEnded() bool {

	return s.End_Of_journey.Before(time.Now())
}
func (s *CacheHarmonize) SetExpirationTime() {
	IsJourney := s.IsJourneyEnded()

	if IsJourney {
		s.isExpired = make(chan bool, 1)
		s.LifeTime = 0
	}
	s.isExpired = make(chan bool, 0)

	s.TimeTTlRemain = calculateTTlTime(s.LifeTime)

}
func calculateTTlTime(value time.Duration) time.Duration {
	//3.9e+6-currentTime  Remain lifetime of  an data
	return time.Duration(65 - value)
}
