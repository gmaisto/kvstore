package encoding

import "github.com/vmihailenco/msgpack/v4"

// MsgPack encodes/decodes Go values to/from Msgpack.
// You can use MsgPack instead of creating an instance of this struct.
type MsgPackCodec struct{}

// Marshal encodes a Go value to Msgpack.
func (c MsgPackCodec) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal decodes a Msgpack value into a Go value.
func (c MsgPackCodec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
