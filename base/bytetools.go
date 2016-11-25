package base

import (
	"bytes"
	"encoding/binary"
	"strconv"
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

var bcd_table [256]string = [256]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "0a", "0b", "0c", "0d", "0e", "0f", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "1a", "1b", "1c", "1d", "1e", "1f", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "2a", "2b", "2c", "2d", "2e", "2f", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "3a", "3b", "3c", "3d", "3e", "3f", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "4a", "4b", "4c", "4d", "4e", "4f", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "5a", "5b", "5c", "5d", "5e", "5f", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "6a", "6b", "6c", "6d", "6e", "6f", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "7a", "7b", "7c", "7d", "7e", "7f", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "8a", "8b", "8c", "8d", "8e", "8f", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "9a", "9b", "9c", "9d", "9e", "9f", "a0", "a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "aa", "ab", "ac", "ad", "ae", "af", "b0", "b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8", "b9", "ba", "bb", "bc", "bd", "be", "bf", "c0", "c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "ca", "cb", "cc", "cd", "ce", "cf", "d0", "d1", "d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "da", "db", "dc", "dd", "de", "df", "e0", "e1", "e2", "e3", "e4", "e5", "e6", "e7", "e8", "e9", "ea", "eb", "ec", "ed", "ee", "ef", "f0", "f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8", "f9", "fa", "fb", "fc", "fd", "fe", "ff"}

func ReadBcdString(reader *bytes.Reader, buffer_count uint8) string {
	bcd_bytes := make([]byte, buffer_count)
	reader.Read(bcd_bytes)

	var result string = ""
	for v := range bcd_bytes {
		result += bcd_table[bcd_bytes[v]]
	}

	return result
}

func ReadBcdTime(reader *bytes.Reader) uint64 {
	bcd_time_string := ReadBcdString(reader, 7)
	bcd_time, _ := strconv.ParseUint(bcd_time_string, 10, 64)

	return bcd_time
}
