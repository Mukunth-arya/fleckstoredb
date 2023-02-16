package fleckstoredb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	Expect = []byte{}

	TimeStamp  = 12345678910
	KeySize    = 12
	ValueSize  = 15
	flags      = 1
	TTL        = 1
	BucketSize = 18
	status     = 0
	Txid       = 0
)

type TestEncodeddata struct {
	suite.Suite
	Test       Entry
	ExpectData []byte
}

func NewInit() *TestEncodeddata {

	return &TestEncodeddata{
		Test: Entry{
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
		ExpectData: []byte{232, 71, 195, 158, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 115, 97, 109, 112, 108, 101, 116, 111, 116, 97, 108, 112, 97, 121, 108, 111, 97, 100, 115, 97, 109, 112, 108, 101, 75, 101, 121, 0, 0, 0, 115, 97, 109, 112, 108, 101, 118, 97, 108, 117, 101, 115, 0, 0, 0},
	}

}
func (s *TestEncodeddata) TestEncodeData() {

	// sample payload size for testing

	Ok_to_process := reflect.DeepEqual(s.Test.Endode_Payload(), s.ExpectData)
	assert.True(s.T(), Ok_to_process, "Encoding Test is failed")

}
func (s *TestEncodeddata) TestNilDataReturn() {
	if oK := s.Test.CheckPayloadisEmpty(); oK {
		assert.Fail(s.T(), "entryIsZeroTestFail")
	}
}

func TestPayload(t *testing.T) {

	suite.Run(t, new(TestEncodeddata))
}
