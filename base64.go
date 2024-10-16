package base64

const mapper = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const symbolZeroByte = "="

func Encode(str string) string {
	var result string

	// Строка в виде байт
	bytes := []byte(str)

	// Строка в виде байт в чанках по 3
	splitBytes := splitBytesTo3ByteChunk(bytes)

	// Длина массива байт с чанками
	lenSplitBytes := len(splitBytes)

	// Количество добавленных 0 байт
	countZeroBytes := lenSplitBytes*3 - len(bytes)

	for i := 0; i < lenSplitBytes; i++ {
		asciiBytes := transform3BytesTo4Bytes(splitBytes[i])

		if countZeroBytes > 0 && i == lenSplitBytes-1 {
			result += convertBytesToString(asciiBytes, countZeroBytes)
		} else {
			result += convertBytesToString(asciiBytes, 0)
		}
	}

	return result
}

// Разбивка массива байт на чанки по 3 байта
func splitBytesTo3ByteChunk(bytes []byte) [][3]byte {
	var result [][3]byte

	j := 0

	var chunk [3]byte

	for i := 0; i < len(bytes); i++ {
		if i%3 == 0 && i != 0 {
			result = append(result, chunk)
			chunk = [3]byte{0, 0, 0}
			j = 0
		}

		chunk[j] = bytes[i]
		j++
	}

	return append(result, chunk)
}

// Конвертация байтов в строку
func convertBytesToString(bytes []byte, countZeroBytes int) (result string) {
	for i := 0; i < len(bytes); i++ {
		// Для добавленных нулевых байт
		if countZeroBytes > 0 && len(bytes)-i <= countZeroBytes {
			result += symbolZeroByte
		} else {
			result += string(mapper[int(bytes[i])])
		}
	}

	return result
}

// Трансформация буффера из 3 байтов в буффер из 4 байтов
func transform3BytesTo4Bytes(bytes [3]byte) []byte {
	b1 := bytes[0]
	b2 := bytes[1]
	b3 := bytes[2]

	result := make([]byte, 4)

	// 1 результирующий байт получаем из 6 старших бит 1 байта
	result[0] = extractBites(b1, 2, 7)

	// 2 результирующий байт получаем путем сложения 2 младших бит 1 байта и 4 старших бит 2 байта
	result[1] = byte(0)
	mask2 := byte(0b00000011)
	result[1] = (b1&mask2)<<4 | b2>>4

	// 3 результирующий байт получаем путем сложения 4 младших бит 2 байта и 2 старших бит 3 байта
	result[2] = byte(0)
	mask3 := byte(0b00001111)
	result[2] = (b2&mask3)<<2 | b3>>6

	// 4 результирующий байт получаем из 6 младших бит 3 байта
	result[3] = extractBites(b3, 0, 5)

	return result
}

// Извлечение диапазона битов из байта
func extractBites(b byte, start int, end int) byte {
	mask := byte(0)

	for i := start; i <= end; i++ {
		mask |= byte(1) << i
	}

	return (b & mask) >> start
}
