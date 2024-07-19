package data

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

type UUID = uuid.UUID

func NewUUID() UUID {
	gen := uuid.NewV4()
	return gen
}

func TextToUUID(text string) (data UUID, err error) {
	return uuid.FromString(text)
}

func BitsToUUID(msb, lsb int64) (data UUID, err error) {
	mBytes := make([]byte, 8)
	lBytes := make([]byte, 8)

	binary.BigEndian.PutUint64(mBytes, uint64(msb))
	binary.BigEndian.PutUint64(lBytes, uint64(lsb))

	return uuid.FromBytes(append(mBytes, lBytes...))
}

func UUIDToText(uuid UUID) (text string, err error) {
	data, err := uuid.MarshalText()

	if err == nil {
		text = string(data)
	}

	return
}

func SigBits(uuid UUID) (msb, lsb int64) {
	bytes := uuid.Bytes()

	msb = 0
	lsb = 0

	for i := 0; i < 8; i++ {
		msb = (msb << 0x08) | int64(bytes[i]&0xFF)
	}

	for i := 8; i < 16; i++ {
		lsb = (lsb << 0x08) | int64(bytes[i]&0xFF)
	}

	return
}

func CreateUUID(line string) []byte {
	if len(line) > 16 {
		return nil
	}

	var buffer bytes.Buffer

	nameInHex := hex.EncodeToString([]byte(line))
	length := len(nameInHex)

	if length <= 12 {
		buffer.WriteString("00000000-0000-4000-0000-")
		buffer.WriteString(nameInHex)
		diference := 12 - length
		for i := 0; i < diference; i++ {
			buffer.WriteRune('0')
		}
		return buffer.Bytes()
	}
	if length <= 20 {
		buffer.WriteString("00000000-4000-")
		buffer.WriteString(nameInHex[:4])
		buffer.WriteRune('-')
		buffer.WriteString(nameInHex[4:8])
		buffer.WriteRune('-')
		buffer.WriteString(nameInHex[8:])

		diference := 36 - buffer.Len()
		for i := 0; i < diference; i++ {
			buffer.WriteRune('0')
		}
		return buffer.Bytes()
	}
	buffer.WriteString(nameInHex[:8])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[8:12])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[12:16])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[16:20])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[20:])

	diference := 36 - buffer.Len()
	for i := 0; i < diference; i++ {
		buffer.WriteRune('0')
	}
	return buffer.Bytes()
}
