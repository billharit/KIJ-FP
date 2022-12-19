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

type SBoxReverse struct {
	data map[byte]byte
}

func (s *SBoxReverse) Init() {
	s.data = make(map[byte]byte)
	s.Generate([]string{"52", "09", "6a", "d5", "30", "36", "a5", "38", "bf", "40", "a3", "9e", "81", "f3", "d7", "fb",
		"7c", "e3", "39", "82", "9b", "2f", "ff", "87", "34", "8e", "43", "44", "c4", "de", "e9", "cb",
		"54", "7b", "94", "32", "a6", "c2", "23", "3d", "ee", "4c", "95", "0b", "42", "fa", "c3", "4e",
		"08", "2e", "a1", "66", "28", "d9", "24", "b2", "76", "5b", "a2", "49", "6d", "8b", "d1", "25",
		"72", "f8", "f6", "64", "86", "68", "98", "16", "d4", "a4", "5c", "cc", "5d", "65", "b6", "92",
		"6c", "70", "48", "50", "fd", "ed", "b9", "da", "5e", "15", "46", "57", "a7", "8d", "9d", "84",
		"90", "d8", "ab", "00", "8c", "bc", "d3", "0a", "f7", "e4", "58", "05", "b8", "b3", "45", "06",
		"d0", "2c", "1e", "8f", "ca", "3f", "0f", "02", "c1", "af", "bd", "03", "01", "13", "8a", "6b",
		"3a", "91", "11", "41", "4f", "67", "dc", "ea", "97", "f2", "cf", "ce", "f0", "b4", "e6", "73",
		"96", "ac", "74", "22", "e7", "ad", "35", "85", "e2", "f9", "37", "e8", "1c", "75", "df", "6e",
		"47", "f1", "1a", "71", "1d", "29", "c5", "89", "6f", "b7", "62", "0e", "aa", "18", "be", "1b",
		"fc", "56", "3e", "4b", "c6", "d2", "79", "20", "9a", "db", "c0", "fe", "78", "cd", "5a", "f4",
		"1f", "dd", "a8", "33", "88", "07", "c7", "31", "b1", "12", "10", "59", "27", "80", "ec", "5f",
		"60", "51", "7f", "a9", "19", "b5", "4a", "0d", "2d", "e5", "7a", "9f", "93", "c9", "9c", "ef",
		"a0", "e0", "3b", "4d", "ae", "2a", "f5", "b0", "c8", "eb", "bb", "3c", "83", "53", "99", "61",
		"17", "2b", "04", "7e", "ba", "77", "d6", "26", "e1", "69", "14", "63", "55", "21", "0c", "7d"})
}

func (s *SBoxReverse) Generate(data []string) {
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

func (m *Matrix) InverseSubtitute() {
	for i, row := range m.data {
		for k, b := range row {
			// log.Printf("Changing %02x -> %02x", b, sbox.data[b])
			m.data[i][k] = sboxReverse.data[b]
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
	return newRow[shiftCount : 4+shiftCount]
}

func (m *Matrix) InverseRoundShift() {
	for i, row := range m.data {
		newRow := m.inverseShiftRow(i, row)
		for k, v := range newRow {
			m.data[i][k] = v
		}
	}
}

func (m *Matrix) inverseShiftRow(shiftCount int, row []byte) []byte {
	newRow := make([]byte, len(row))
	copy(newRow, row)
	newRow = append(newRow, newRow...)
	return newRow[len(row)-shiftCount : 4+(len(row)-shiftCount)]
}

func (m *Matrix) MixColumn() {
	fixedMatrix := [][]byte{
		{2, 3, 1, 1},
		{1, 2, 3, 1},
		{1, 1, 2, 3},
		{3, 1, 1, 2},
	}

	multiply := func(a byte, b byte) byte {
		// a7 = bit 7 from a
		// b0 = bit 0 from b
		var p byte
		r := byte(1<<4 | 1<<3 | 1<<1 | 1<<0)
		// r := byte(1<<4 | 1<<0)
		for i := 0; i < 8; i++ {
			ab0 := a &^ (b&1 - 1)  // ab0 = Mul(a, b0)
			ra7 := r &^ (a>>7 - 1) // ra7 = Mul(R, a7)

			// p = Add(p, ab0)
			// a = Sub(Mul(a, x), ra7)
			// b = Div(b, x)
			p, a, b = p^ab0, a<<1^ra7, b>>1
		}
		return p
	}

	newRow := [][]byte{}
	n := 0
	for n < 4 {
		newRow = append(newRow, []byte{})
		n++
	}
	for i, row := range m.data {
		nowRow := []byte{}
		for k := range row {
			var newValue byte
			for n := 0; n < 4; n++ {
				if n == 0 {
					newValue = multiply(fixedMatrix[i][n], m.data[n][k])
				} else {
					newValue = newValue ^ multiply(fixedMatrix[i][n], m.data[n][k])
				}
			}
			nowRow = append(nowRow, newValue)
		}
		newRow[i] = nowRow
	}
	m.data = newRow
}

func (m *Matrix) InverseMixColumn() {
	fixedMatrix := [][]byte{
		{14, 11, 13, 9},
		{9, 14, 11, 13},
		{13, 9, 14, 11},
		{11, 13, 9, 14},
	}

	multiply := func(a byte, b byte) byte {
		// a7 = bit 7 from a
		// b0 = bit 0 from b
		var p byte
		r := byte(1<<4 | 1<<3 | 1<<1 | 1<<0)
		for i := 0; i < 8; i++ {
			ab0 := a &^ (b&1 - 1)  // ab0 = Mul(a, b0)
			ra7 := r &^ (a>>7 - 1) // ra7 = Mul(R, a7)

			// p = Add(p, ab0)
			// a = Sub(Mul(a, x), ra7)
			// b = Div(b, x)
			p, a, b = p^ab0, a<<1^ra7, b>>1
		}
		return p
	}

	// region - init new state matrix
	newRow := [][]byte{}
	n := 0
	for n < 4 {
		newRow = append(newRow, []byte{})
		n++
	}
	// region - init new state matrix

	for i, row := range m.data {
		nowRow := []byte{}
		for k := range row {
			var newValue byte
			for n := 0; n < 4; n++ {
				if n == 0 {
					newValue = multiply(fixedMatrix[i][n], m.data[n][k])
				} else {
					newValue = newValue ^ multiply(fixedMatrix[i][n], m.data[n][k])
				}
			}
			nowRow = append(nowRow, newValue)
		}
		newRow[i] = nowRow
	}
	m.data = newRow
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
	fmt.Printf("===========\n")
}

func (m *Matrix) PrintString() {
	fmt.Printf(" === ")
	for i := 0; i < len(m.data); i++ {
		for j := 0; j < len(m.data); j++ {
			fmt.Printf("%02x", m.data[j][i])
		}
		fmt.Printf(" ")
	}
	fmt.Printf("\n")
}

func (m *Matrix) String() string {
	result := ""
	for i := 0; i < len(m.data); i++ {
		for j := 0; j < len(m.data); j++ {
			result += fmt.Sprintf("%02x", m.data[j][i])
		}
		result += fmt.Sprintf(" ")
	}
	return result
}

// func Pretty(input []byte) string {
// 	result := ""
// 	for i := 0; i < len(input); i++ {
// 		for j := 0; j < len(input); j++ {
// 			result += fmt.Sprintf("%02x", m.data[j][i])
// 		}
// 		result += fmt.Sprintf(" ")
// 	}
// 	return result
// }

func (m *Matrix) AsByte() []byte {
	result := []byte{}
	for i := 0; i < len(m.data); i++ {
		for j := 0; j < len(m.data); j++ {
			result = append(result, m.data[j][i])
		}
	}
	return result
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
var sboxReverse *SBoxReverse

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

func shiftRow(shiftCount uint8, row []byte) []byte {
	row = append(row, row...)
	return row[shiftCount : 3+shiftCount]
}

func encryptECBBlock(plainText []byte, key []byte) []byte {
	// region - Rjindael Sbox
	sboxLocal := &SBox{}
	sboxLocal.Init()
	sbox = sboxLocal

	sboxReverseLocal := &SBoxReverse{}
	sboxReverseLocal.Init()
	sboxReverse = sboxReverseLocal
	// region - Rjindael Sbox

	totalRound := 10

	// region - Compute Round Constant for Key Expansion
	preComputeRoundConstant(totalRound)
	// region - Compute Round Constant for Key Expansion

	// region - Key Expansion
	words := [][]byte{}

	words = append(words, key[0:4])
	words = append(words, key[4:8])
	words = append(words, key[8:12])
	words = append(words, key[12:16])
	log.Println("Key to byte", hex.EncodeToString(key), len(key))

	roundConstant := 0
	for {
		// docs - Expand key for 4*totalRound times
		if len(words) == (totalRound+1)*4 {
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

	// region - Key Expansion

	// region - Prepare plaintext as matrix
	matrix := &Matrix{}
	matrix.Init()
	matrix.Fill(plainText)
	// region - Prepare plaintext as matrix

	// region - Encryption
	round := 0
	for round <= totalRound {
		if round > 0 {
			matrix.Subtitute()
			if debugMode {
				fmt.Scanln(&buffer)
			}
			log.Println("After subtitutes", round)
			matrix.Print()

			matrix.RoundShift()
			if debugMode {
				fmt.Printf("Next shift")
				fmt.Scanln(&buffer)
			}
			log.Println("After shift", round)
			matrix.Print()

			if round < totalRound {
				if debugMode {
					fmt.Printf("Next mixCol")
					fmt.Scanln(&buffer)
				}
				matrix.MixColumn()
				log.Println("After mixCol", round)
				matrix.Print()
			}
		}

		key = joinMatrix(words[(round * 4) : (round*4)+4]...)
		keyMatrix := &Matrix{}
		keyMatrix.Init()
		keyMatrix.Fill(key)

		matrix.XOR(keyMatrix)
		if debugMode {
			fmt.Scanln(&buffer)
		}
		log.Println("After round key", round)
		matrix.Print()

		round++
	}
	// region - Encryption

	return matrix.AsByte()
}

var buffer = ""

func decryptECBBlock(cipherText []byte, key []byte) []byte {
	// region - Algo Properties
	round := 0
	totalRound := 10
	decryptRound := 10
	// region - Algo Properties

	// region - Key Expansion
	words := [][]byte{}

	words = append(words, key[0:4])
	words = append(words, key[4:8])
	words = append(words, key[8:12])
	words = append(words, key[12:16])

	roundConstant := 0
	for {
		// docs - Expand key for 4*totalRound times
		if len(words) == (totalRound+1)*4 {
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
	// region - Key Expansion

	// region - Decryption
	matrix := &Matrix{}
	matrix.Init()
	matrix.Fill(cipherText)
	for round <= totalRound {
		if round > 0 {
			matrix.InverseRoundShift()
			// log.Println("Decrypt after shift", round)
			// matrix.Print()
			// matrix.PrintString()

			matrix.InverseSubtitute()
			// log.Println("Decrypt after subtitute", round)
			// matrix.Print()
			// matrix.PrintString()

			if decryptRound >= 0 {
				key = joinMatrix(words[(decryptRound * 4) : (decryptRound*4)+4]...)
				keyMatrix := &Matrix{}
				keyMatrix.Init()
				keyMatrix.Fill(key)
				// log.Println("Decrypt AddRoundKey", decryptRound)
				// keyMatrix.PrintString()

				matrix.XOR(keyMatrix)
				// matrix.Print()
				// matrix.PrintString()

				if decryptRound > 0 {
					matrix.InverseMixColumn()
					// log.Println("Decrypt after mixCol", round)
					// matrix.Print()
					// matrix.PrintString()
				}
			}
		}

		if round == 0 {
			key = joinMatrix(words[(decryptRound * 4) : (decryptRound*4)+4]...)
			keyMatrix := &Matrix{}
			keyMatrix.Init()
			keyMatrix.Fill(key)
			// log.Println("Decrypt AddRoundKey", decryptRound)

			matrix.XOR(keyMatrix)
			// matrix.Print()
			// matrix.PrintString()
		}
		round++
		decryptRound--
	}

	return matrix.AsByte()
}

func encryptECB(input []byte, key []byte) []byte {
	result := []byte{}
	iteration := 0

	// pad input
	for len(input)%16 > 0 {
		input = append(input, byte(0))
	}

	log.Println("Input len after pad", len(input), string(input))

	for {
		block := input[iteration*16 : (iteration+1)*16]
		encrpytedBlock := encryptECBBlock(block, key)
		result = append(result, encrpytedBlock...)
		iteration++
		if len(input) == (iteration)*16 {
			break
		}
	}
	return result
}

func decryptECB(input []byte, key []byte) []byte {
	result := []byte{}
	iteration := 0
	for {
		block := decryptECBBlock(input[iteration*16:(iteration+1)*16], key)
		result = append(result, block...)
		iteration++
		if len(input) == (iteration)*16 {
			break
		}
	}
	return result
}

var debugMode bool = false

func main() {
	plainText := "Two One Nine TwoTwo One Nine Tw"
	key := "1234123412341234"
	fmt.Println("Plain text:", plainText, hex.EncodeToString([]byte(plainText)))
	fmt.Println("Key:", key, hex.EncodeToString([]byte(key)))

	encryptedCBC := encryptECB([]byte(plainText), []byte(key))
	log.Println("Encrypted ECB (len)", len(encryptedCBC), hex.EncodeToString(encryptedCBC))
	decryptedCBC := decryptECB(encryptedCBC, []byte(key))
	fmt.Println("Plaintext after decrypt ECB (byte):", decryptedCBC)
	fmt.Println("Plaintext after decrypt ECB (hex):", hex.EncodeToString(decryptedCBC))
	fmt.Println("Plaintext after decrypt ECB:", string(decryptedCBC))
}
