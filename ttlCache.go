package fleckstoredb

import "time"

const (
	// Default TTl DataDuration
	// Default amount Of That the data is present in Voltile Memory
	DefaultTTLTime time.Duration = 20
	//TTl Gona Never Expire
	TTL_NeverExpire time.Duration = -10
)

type CacheHarmonize struct {
	EntryData      *Entry
	LifeTime       time.Duration
	End_Of_journey time.Time
	Stuck_here     bool
}

type TTlCache struct {
}
