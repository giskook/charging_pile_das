package base

import (
	"bytes"
	"encoding/binary"
)

func ReadWord(reader *bytes.Reader) uint16 {
	word_byte := make([]byte, 2)
	reader.Read(word_byte)

	return binary.BigEndian.Uint16(word_byte)
}

func ReadDWord(reader *bytes.Reader) uint32 {
	dword_byte := make([]byte, 4)
	reader.Read(dword_byte)

	return binary.BigEndian.Uint32(dword_byte)
}

func ReadQuaWord(reader *bytes.Reader) uint64 {
	qword_byte := make([]byte, 8)
	reader.Read(qword_byte)

	return binary.BigEndian.Uint64(qword_byte)
}

func ReadMac(reader *bytes.Reader) uint64 {
	mac_byte := make([]byte, 6)
	reader.Read(mac_byte)
	mac := []byte{0, 0}
	mac = append(mac, mac_byte...)

	return binary.BigEndian.Uint64(mac)
}

func ReadString(reader *bytes.Reader, length uint8) string {
	string_byte := make([]byte, length)
	reader.Read(string_byte)

	return string(string_byte)
}

func WriteMac(mac uint64) []byte {
	mac_byte := make([]byte, 8)
	binary.BigEndian.PutUint64(mac_byte, mac)

	return mac_byte[2:]
}

func WriteMacBytes(writer *bytes.Buffer, mac uint64) {
	writer.Write(WriteMac(mac))
}

func WriteWord(writer *bytes.Buffer, word uint16) {
	word_byte := make([]byte, 2)
	binary.BigEndian.PutUint16(word_byte, word)

	writer.Write(word_byte)
}

func WriteDWord(writer *bytes.Buffer, dword uint32) {
	dword_byte := make([]byte, 4)
	binary.BigEndian.PutUint32(dword_byte, dword)

	writer.Write(dword_byte)
}

func WriteQuaWord(writer *bytes.Buffer, quaword uint64) {
	quaword_byte := make([]byte, 8)
	binary.BigEndian.PutUint64(quaword_byte, quaword)

	writer.Write(quaword_byte)
}

func WriteLength(writer *bytes.Buffer) {
	length := writer.Len()
	length += 2
	length_byte := make([]byte, 2)
	binary.BigEndian.PutUint16(length_byte, uint16(length))
	writer.Bytes()[1] = length_byte[0]
	writer.Bytes()[2] = length_byte[1]
}

func GetWord(buffer []byte) uint16 {
	temp := make([]byte, 2)
	temp[0] = buffer[0]
	temp[1] = buffer[1]

	return binary.BigEndian.Uint16(temp)
}
