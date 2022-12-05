package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

type SBox struct {
	data map[byte]byte
}

func (s *SBox) Init() {
	s.data = make(map[byte]byte)
	s.Generate([]string{"63", "7c", "77", "7b", "f2", "6b", "6f", "c5", "30", "01", "67", "2b", "fe", "d7", "ab", "76",
		"ca", "82", "c9", "7d", "fa", "59", "47", "f0", "ad", "d4", "a2", "af", "9c", "a4", "72", "c0",
		"b7", "fd", "93", "26", "36", "3f", "f7", "cc", "34", "a5", "e5", "f1", "71", "d8", "31", "15",
		"04", "c7", "23", "c3", "18", "96", "05", "9a", "07", "12", "80", "e2", "eb", "27", "b2", "75",
		"09", "83", "2c", "1a", "1b", "6e", "5a", "a0", "52", "3b", "d6", "b3", "29", "e3", "2f", "84",
		"53", "d1", "00", "ed", "20", "fc", "b1", "5b", "6a", "cb", "be", "39", "4a", "4c", "58", "cf",
		"d0", "ef", "aa", "fb", "43", "4d", "33", "85", "45", "f9", "02", "7f", "50", "3c", "9f", "a8",
		"51", "a3", "40", "8f", "92", "9d", "38", "f5", "bc", "b6", "da", "21", "10", "ff", "f3", "d2",
		"cd", "0c", "13", "ec", "5f", "97", "44", "17", "c4", "a7", "7e", "3d", "64", "5d", "19", "73",
		"60", "81", "4f", "dc", "22", "2a", "90", "88", "46", "ee", "b8", "14", "de", "5e", "0b", "db",
		"e0", "32", "3a", "0a", "49", "06", "24", "5c", "c2", "d3", "ac", "62", "91", "95", "e4", "79",
		"e7", "c8", "37", "6d", "8d", "d5", "4e", "a9", "6c", "56", "f4", "ea", "65", "7a", "ae", "08",
		"ba", "78", "25", "2e", "1c", "a6", "b4", "c6", "e8", "dd", "74", "1f", "4b", "bd", "8b", "8a",
		"70", "3e", "b5", "66", "48", "03", "f6", "0e", "61", "35", "57", "b9", "86", "c1", "1d", "9e",
		"e1", "f8", "98", "11", "69", "d9", "8e", "94", "9b", "1e", "87", "e9", "ce", "55", "28", "df",
		"8c", "a1", "89", "0d", "bf", "e6", "42", "68", "41", "99", "2d", "0f", "b0", "54", "bb", "16"})
}

func (s *SBox) Generate(data []string) {
	for i, v := range data {
		hexVal, err := hex.DecodeString(v)
		if err != nil {
			panic(err)
		}
		s.data[byte(i)] = hexVal[0]
	}
}

type Matrix struct {
	data [][]byte
}

func (m *Matrix) Init() {
	m.data = make([][]byte, 0)
	n := 0
	for n < 4 {
		m.data = append(m.data, []byte{})
		n++
	}
}

func (m *Matrix) Fill(data []byte) {
	x := 0
	y := 0
	for _, v := range data {
		m.data[x] = append(m.data[x], v)
		x = (x + 1) % 4
		if x == 0 {
			y = (y + 1) % 4
		}
	}
}

func (m *Matrix) XOR(matrix *Matrix) {
	for i, row := range m.data {
		for k, b := range row {
			m.data[i][k] = b ^ matrix.data[i][k]
		}
	}
}

func (m *Matrix) Subtitute() {
	for i, row := range m.data {
		for k, b := range row {
			// log.Printf("Changing %02x -> %02x", b, sbox.data[b])
			m.data[i][k] = sbox.data[b]
		}
	}
}

func (m *Matrix) RoundShift() {
	for i, row := range m.data {
		newRow := m.shiftRow(i, row)
		for k, v := range newRow {
			m.data[i][k] = v
		}
	}
}

func (m *Matrix) shiftRow(shiftCount int, row []byte) []byte {
	newRow := make([]byte, len(row))
	copy(newRow, row)
	newRow = append(newRow, newRow...)
	// log.Println(newRow)
	return newRow[shiftCount : 4+shiftCount]
}

func (m *Matrix) MixColumn() {
	fixedMatrix := [][]byte{
		{2, 3, 1, 1},
		{1, 2, 3, 1},
		{1, 1, 2, 3},
		{3, 1, 1, 2},
	}
	for i, row := range m.data {
		for k, b := range row {
			rowMultiply := fixedMatrix[i]
			columnMultiply := []byte{}
			n := 0
			for n < 4 {
				columnMultiply = append(columnMultiply, m.data[n][k])
				n++
			}
			log.Println(i, k, rowMultiply, columnMultiply)
			newValue := b * rowMultiply[i]
			log.Printf("%02x %02x %b", b, rowMultiply[i], newValue)
			for c, v := range columnMultiply {
				if v == b {
					continue
				}
				log.Printf("%02x %02x %b %d", v, rowMultiply[c], (v * rowMultiply[c]), (v * rowMultiply[c]))
				newValue = newValue ^ (v * rowMultiply[c])
			}
			log.Printf("%02x\n", newValue)

			m.data[i][k] = newValue
		}
	}
}

func (m *Matrix) Print() {
	fmt.Printf("===========\n")
	for _, row := range m.data {
		if len(row) == 0 {
			return
		}
		for _, b := range row {
			fmt.Printf("%02x ", b)
		}
		fmt.Printf("\n")
	}
}

func gFunc(constant byte, key []byte) []byte {
	var result = make([]byte, len(key))
	copy(result, key)
	temp := result[0]

	result[0] = sbox.data[result[1]] ^ (byte(roundConstant[constant]))
	result[1] = sbox.data[result[2]]
	result[2] = sbox.data[result[3]]
	result[3] = sbox.data[temp]

	return result
}

var roundConstant = []byte{}

func matrixXor(a []byte, b []byte) []byte {
	result := []byte{}
	for i, v := range a {
		result = append(result, v^b[i])
	}
	return result
}

var sbox *SBox

func preComputeRoundConstant(n int) {
	i := 0
	for {
		if i == n {
			break
		}
		last := len(roundConstant) - 1
		if i == 0 {
			roundConstant = append(roundConstant, 1)
		} else if roundConstant[last] < 128 {
			roundConstant = append(roundConstant, 2*roundConstant[last])
		} else {
			before := roundConstant[last]
			now := (2 * int(before)) ^ 283
			roundConstant = append(roundConstant, byte(now))
		}
		i++
	}
}

func joinMatrix(words ...[]byte) []byte {
	buffer := []byte{}
	for _, b := range words {
		if len(buffer) == 16 {
			return buffer
		}
		for _, bc := range b {
			buffer = append(buffer, bc)
		}
	}
	return buffer
}

func subtituteOperation(matrix []byte) {
	for i, v := range matrix {
		log.Printf("Changing %02x -> %02x", v, sbox.data[v])
		matrix[i] = sbox.data[v]
	}
}

func shiftMatrix(matrix []byte) {
	for i, _ := range matrix {
		if i%4 == 0 {
			shiftCount := i / 4
			if shiftCount > 0 {
				shiftedRow := shiftRow(uint8(shiftCount), matrix[i:i+4])
				for k, v := range shiftedRow {
					matrix[i+k] = v
				}
			}
		}
	}
}

// func mixColumn(matrix []byte) {
// 	for i, v := range matrix {

// 	}
// }

func shiftRow(shiftCount uint8, row []byte) []byte {
	row = append(row, row...)
	return row[shiftCount : 3+shiftCount]
}

func main() {
	sboxLocal := &SBox{}
	sboxLocal.Init()
	sbox = sboxLocal

	totalRound := 2

	preComputeRoundConstant(2 * totalRound)

	plainText := []byte("Two One Nine Two")
	key := []byte("Thats my Kung Fu")
	log.Println("Key size:", len(key)*8)

	encodedString := hex.EncodeToString(plainText)
	log.Println(encodedString)
	encodedString = hex.EncodeToString(key)
	log.Println(encodedString)

	words := [][]byte{}

	words = append(words, key[0:4])
	words = append(words, key[4:8])
	words = append(words, key[8:12])
	words = append(words, key[12:16])

	roundConstant := 0
	for {
		if len(words) == totalRound*4 {
			break
		}
		last := len(words) - 1
		if len(words)%4 == 0 {
			gMatrix := gFunc(byte(roundConstant), words[last])
			words = append(words, matrixXor(words[last-3], gMatrix)) // word 4
			roundConstant++
		} else {
			words = append(words, matrixXor(words[last], words[last-3]))
		}
	}
	log.Println(len(words))
	buffer := []byte{}
	for i, b := range words {
		if i%4 == 0 {
			fmt.Printf("Round: %d ", (i/4)-1)
			for _, bf := range buffer {
				fmt.Printf("%02x ", bf)
			}
			fmt.Println("")
			buffer = []byte{}
		}
		for _, bc := range b {
			buffer = append(buffer, bc)
		}
	}
	log.Println("Finished")

	matrix := &Matrix{}
	matrix.Init()
	matrix.Fill(plainText)
	matrix.Print()
	round := 0
	for {
		if round == totalRound {
			break
		}
		if round > 0 {
			matrix.Subtitute()
			log.Println("After subtitute", round)
			matrix.Print()

			matrix.RoundShift()
			log.Println("After shift", round)
			matrix.Print()

			matrix.MixColumn()
			log.Println("After mixCol", round)
			matrix.Print()
		}
		key = joinMatrix(words[(round * 4) : (round*4)+4]...)
		keyMatrix := &Matrix{}
		keyMatrix.Init()
		keyMatrix.Fill(key)

		matrix.XOR(keyMatrix)
		log.Println("XOR", round)

		matrix.Print()

		round++
	}
}

// 54776f204f6e65204e696e652054776f
// 5468617473206d79204b756e67204675
