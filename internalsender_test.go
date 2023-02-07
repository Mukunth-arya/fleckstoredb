package fleckstoredb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Expect           = []byte{}
	TotalPayloadSize uint32
)

type TestEncodeddata struct {
	Test *Entry
}

func NewInit() *TestEncodeddata {

	return &TestEncodeddata{
		Test: &Entry{
			Keys:  []byte("sampleKey"),
			Value: []byte("samplevalues"),
			Meta: &Metadata{
				KeySize:   uint32(len("sampleKey001")),
				ValueSize: uint32(len("samplevalues002")),
				TimeStamp: 12345678910,
				// If ttl is setted to be 1 it is persistent
				// rather than it is setted to be an exact time
				// The exact the data will be deleted
				TTL: 1,
				// If flag is 1 the data is setted
				// Rather than 1 like 0 the data will be not setted
				flags:      1,
				Bucket:     []byte("sampletotalpayload"),
				BucketSize: uint32(len("sampletotalpayload")),
			},
		},
	}

}
func SetTheExactPayLoadSize() {
	if len(Expect) == 0 {

		Expect = []byte{232, 71, 195, 158, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 5, 97, 109, 112, 108, 101, 116, 111, 116, 97, 108, 112, 97, 121, 108, 111, 97, 100, 115, 97, 109, 112, 108, 101, 75, 101, 121, 0, 0, 0, 115, 97, 109, 112, 108, 101, 118, 97, 108, 117, 101, 115, 0, 0, 0}
	}
}
func (s *TestEncodeddata) TestencodeData(t *testing.T) {
	SetTheExactPayLoadSize()
	// sample payload size for testing
	TotalPayloadSize = 50
	Ok_to_process := reflect.DeepEqual(s.Test.Endode_Payload(), Expect)
	assert.True(t, Ok_to_process, "Encoding is failed")

}
