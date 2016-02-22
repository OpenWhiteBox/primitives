package encoding

// ParseByte parses a serialized Byte encoding.
func ParseByte(serialized []byte) Byte {
	sbox := SBox{}

	for in := 0; in < 256; in++ {
		sbox.EncKey[in] = serialized[in]
		sbox.DecKey[serialized[in]] = byte(in)
	}

	return sbox
}

// SerializeByte serializes a Byte encoding to a byte slice.
func SerializeByte(e Byte) []byte {
	out := make([]byte, 256)
	for in := 0; in < 256; in++ {
		out[in] = e.Encode(byte(in))
	}

	return out
}
