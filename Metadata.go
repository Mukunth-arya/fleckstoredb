package fleckstoredb

type (
	Metadata struct {
		KeySize    uint32
		ValueSize  uint32
		TimeStamp  uint64
		TTL        uint32
		flags      uint16
		Indicator  uint16
		TXid       uint16
		Bucket     []byte
		BucketSize uint32
		Status     uint16
	}
)
