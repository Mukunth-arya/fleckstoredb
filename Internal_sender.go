package fleckstoredb

import (
	"encoding/binary"
	"hash/crc32"
)

type Entry struct {
	Keys  []byte
	Value []byte
	Meta  *Metadata
	Crc   uint32
}

// The indexes is the giving hint to search key
type Indexes struct {
	Key    []byte
	FileId int64
	Meta   *Metadata
}

//Following function endcode the payload
//The encoded payload is stored in the format of Little endian
//EX :12 ------> Encoded format is 1->49,2->50
//In the Ram these two encoded payload is stored in the  format like:
// _________________________
//|           |             |
//|     50    |     49      |
//|           |             |
//|___________|_____________|

//Making The Payload  Size entry to be stored
func (e *Entry) TotalPayloadStoreSize() uint32 {

	return (TotalPayloadSize + e.Meta.KeySize + e.Meta.ValueSize + e.Meta.BucketSize)
}
func (e *Entry) Endode_Payload() []byte {

	payload := make([]byte, e.TotalPayloadStoreSize())

	// copy the remaing data rather than the metadata  and
	copy(payload[(TotalPayloadSize):(TotalPayloadSize+e.Meta.BucketSize)], e.Meta.Bucket)
	copy(payload[(TotalPayloadSize+e.Meta.BucketSize):(TotalPayloadSize+e.Meta.BucketSize+e.Meta.KeySize)], e.Keys)
	copy(payload[(TotalPayloadSize+e.Meta.BucketSize+e.Meta.KeySize):(TotalPayloadSize+e.Meta.BucketSize+e.Meta.KeySize+e.Meta.ValueSize)], e.Value)

	// The payload is filled with the crc bits and an data
	// crc means the cyclic redunancy check
	// If remainder is 00000 -> Data is not coruppted
	// If the remainder is neither the 0000 and mixed with the 010101
	// it will  not be an zero
	// _________________________
	//|           |             |
	//|     CRC   |     Data    |
	//|     BITS  |     BITS    |
	//|___________|_____________|
	// The Bytes of slice of data is given to crc
	// The Remainder data 0 will be negleted and remaining will be crc bits
	crc := crc32.ChecksumIEEE(payload[4:])
	binary.LittleEndian.PutUint32(payload[0:4], crc)
	return payload
}

// Make payload stored in the little indian format
func (e *Entry) MakePayloadEntry(payload []byte) {
	binary.LittleEndian.PutUint64(payload[4:12], e.Meta.TimeStamp)
	binary.LittleEndian.PutUint32(payload[12:16], uint32(e.Meta.KeySize))
	binary.LittleEndian.PutUint32(payload[16:20], uint32(e.Meta.ValueSize))
	binary.LittleEndian.PutUint16(payload[20:22], uint16(e.Meta.flags))
	binary.LittleEndian.PutUint32(payload[22:26], uint32(e.Meta.TTL))
	binary.LittleEndian.PutUint32(payload[26:30], uint32(e.Meta.BucketSize))

	binary.LittleEndian.PutUint16(payload[32:34], uint16(e.Meta.Status))
	binary.LittleEndian.PutUint16(payload[34:36], uint16(e.Meta.TXid))

}
