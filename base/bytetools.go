package base

import (
	"bytes"
	"encoding/binary"
	"log"
	"strconv"
	"time"
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

func WriteString(writer *bytes.Buffer, str string) {
	writer.Write([]byte(str))
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
	length += 3
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
var bcd_hex2str_map map[uint8]string = map[uint8]string{0x00: "00", 0x01: "01", 0x02: "02", 0x03: "03", 0x04: "04", 0x05: "05", 0x06: "06", 0x07: "07", 0x08: "08", 0x09: "09", 0x10: "10", 0x11: "11", 0x12: "12", 0x13: "13", 0x14: "14", 0x15: "15", 0x16: "16", 0x17: "17", 0x18: "18", 0x19: "19", 0x20: "20", 0x21: "21", 0x22: "22", 0x23: "23", 0x24: "24", 0x25: "25", 0x26: "26", 0x27: "27", 0x28: "28", 0x29: "29", 0x30: "30", 0x31: "31", 0x32: "32", 0x33: "33", 0x34: "34", 0x35: "35", 0x36: "36", 0x37: "37", 0x38: "38", 0x39: "39", 0x40: "40", 0x41: "41", 0x42: "42", 0x43: "43", 0x44: "44", 0x45: "45", 0x46: "46", 0x47: "47", 0x48: "48", 0x49: "49", 0x50: "50", 0x51: "51", 0x52: "52", 0x53: "53", 0x54: "54", 0x55: "55", 0x56: "56", 0x57: "57", 0x58: "58", 0x59: "59", 0x60: "60", 0x61: "61", 0x62: "62", 0x63: "63", 0x64: "64", 0x65: "65", 0x66: "66", 0x67: "67", 0x68: "68", 0x69: "69", 0x70: "70", 0x71: "71", 0x72: "72", 0x73: "73", 0x74: "74", 0x75: "75", 0x76: "76", 0x77: "77", 0x78: "78", 0x79: "79", 0x80: "80", 0x81: "81", 0x82: "82", 0x83: "83", 0x84: "84", 0x85: "85", 0x86: "86", 0x87: "87", 0x88: "88", 0x89: "89", 0x90: "90", 0x91: "91", 0x92: "92", 0x93: "93", 0x94: "94", 0x95: "95", 0x96: "96", 0x97: "97", 0x98: "98", 0x99: "99"}
var bcd_map map[string]uint8 = map[string]uint8{"00": 0, "01": 1, "02": 2, "03": 3, "04": 4, "05": 5, "06": 6, "07": 7, "08": 8, "09": 9, "0a": 10, "0b": 11, "0c": 12, "0d": 13, "0e": 14, "0f": 15, "10": 16, "11": 17, "12": 18, "13": 19, "14": 20, "15": 21, "16": 22, "17": 23, "18": 24, "19": 25, "1a": 26, "1b": 27, "1c": 28, "1d": 29, "1e": 30, "1f": 31, "20": 32, "21": 33, "22": 34, "23": 35, "24": 36, "25": 37, "26": 38, "27": 39, "28": 40, "29": 41, "2a": 42, "2b": 43, "2c": 44, "2d": 45, "2e": 46, "2f": 47, "30": 48, "31": 49, "32": 50, "33": 51, "34": 52, "35": 53, "36": 54, "37": 55, "38": 56, "39": 57, "3a": 58, "3b": 59, "3c": 60, "3d": 61, "3e": 62, "3f": 63, "40": 64, "41": 65, "42": 66, "43": 67, "44": 68, "45": 69, "46": 70, "47": 71, "48": 72, "49": 73, "4a": 74, "4b": 75, "4c": 76, "4d": 77, "4e": 78, "4f": 79, "50": 80, "51": 81, "52": 82, "53": 83, "54": 84, "55": 85, "56": 86, "57": 87, "58": 88, "59": 89, "5a": 90, "5b": 91, "5c": 92, "5d": 93, "5e": 94, "5f": 95, "60": 96, "61": 97, "62": 98, "63": 99, "64": 100, "65": 101, "66": 102, "67": 103, "68": 104, "69": 105, "6a": 106, "6b": 107, "6c": 108, "6d": 109, "6e": 110, "6f": 111, "70": 112, "71": 113, "72": 114, "73": 115, "74": 116, "75": 117, "76": 118, "77": 119, "78": 120, "79": 121, "7a": 122, "7b": 123, "7c": 124, "7d": 125, "7e": 126, "7f": 127, "80": 128, "81": 129, "82": 130, "83": 131, "84": 132, "85": 133, "86": 134, "87": 135, "88": 136, "89": 137, "8a": 138, "8b": 139, "8c": 140, "8d": 141, "8e": 142, "8f": 143, "90": 144, "91": 145, "92": 146, "93": 147, "94": 148, "95": 149, "96": 150, "97": 151, "98": 152, "99": 153, "9a": 154, "9b": 155, "9c": 156, "9d": 157, "9e": 158, "9f": 159, "a0": 160, "a1": 161, "a2": 162, "a3": 163, "a4": 164, "a5": 165, "a6": 166, "a7": 167, "a8": 168, "a9": 169, "aa": 170, "ab": 171, "ac": 172, "ad": 173, "ae": 174, "af": 175, "b0": 176, "b1": 177, "b2": 178, "b3": 179, "b4": 180, "b5": 181, "b6": 182, "b7": 183, "b8": 184, "b9": 185, "ba": 186, "bb": 187, "bc": 188, "bd": 189, "be": 190, "bf": 191, "c0": 192, "c1": 193, "c2": 194, "c3": 195, "c4": 196, "c5": 197, "c6": 198, "c7": 199, "c8": 200, "c9": 201, "ca": 202, "cb": 203, "cc": 204, "cd": 205, "ce": 206, "cf": 207, "d0": 208, "d1": 209, "d2": 210, "d3": 211, "d4": 212, "d5": 213, "d6": 214, "d7": 215, "d8": 216, "d9": 217, "da": 218, "db": 219, "dc": 220, "dd": 221, "de": 222, "df": 223, "e0": 224, "e1": 225, "e2": 226, "e3": 227, "e4": 228, "e5": 229, "e6": 230, "e7": 231, "e8": 232, "e9": 233, "ea": 234, "eb": 235, "ec": 236, "ed": 237, "ee": 238, "ef": 239, "f0": 240, "f1": 241, "f2": 242, "f3": 243, "f4": 244, "f5": 245, "f6": 246, "f7": 247, "f8": 248, "f9": 249, "fa": 250, "fb": 251, "fc": 252, "fd": 253, "fe": 254, "ff": 255}

func ReadBcdString(reader *bytes.Reader, buffer_count uint8) string {
	bcd_bytes := make([]byte, buffer_count)
	reader.Read(bcd_bytes)

	var result string = ""
	for v := range bcd_bytes {
		result += bcd_table[bcd_bytes[v]]
	}

	return result
}

func WriteBcdString(writer *bytes.Buffer, str string) {
	str_len := len(str)
	for i := 0; i < str_len/2; i++ {
		writer.WriteByte(bcd_map[str[i*2:i*2+2]])
	}
}

func ReadBcdTime(reader *bytes.Reader) uint64 {
	bcd_time_string := ReadBcdString(reader, 6)
	//loc, _ := time.LoadLocation("Asia/Beijing")

	_time, _ := time.ParseInLocation("20060102150405", "20"+bcd_time_string, time.Local)
	log.Println("---------")
	log.Println(bcd_time_string)
	log.Println(_time.Unix())

	return uint64(_time.Unix())
}

func WriteBcdCpid(writer *bytes.Buffer, cpid uint64) {
	cpid_str := strconv.FormatUint(cpid, 10)
	WriteBcdString(writer, cpid_str)
}
