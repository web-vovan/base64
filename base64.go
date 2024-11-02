package base64

import (
	"strings"
)

const mapper = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const symbolZeroByte = "="

// Encode Кодирование строки в base64
func Encode(str string) string {
	var result string

	// Строка в виде байт
	bytes := []byte(str)

	// Строка в виде байт в чанках по 3
	splitBytes := splitBytesToByteChunk(bytes, 3)

	// Длина массива байт с чанками
	lenSplitBytes := len(splitBytes)

	// Количество добавленных нулевых байт
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

// Decode Декодирование base64 в строку
func Decode(str string) string {
	var result string

	// Удаляем символы нулевого байта из конца строки
	str = trimSymbolZeroByte(str)

	// Слайс байт с индексами mapper
	var indexBytes []byte

	for i := 0; i < len(str); i++ {
		// Индекс строки из mapper
		index := strings.Index(mapper, string(str[i]))
		indexBytes = append(indexBytes, byte(index))
	}

	// Разбиваем байты на чанки по 4
	splitBytes := splitBytesToByteChunk(indexBytes, 4)

	// Длина массива байт с чанками
	lenSplitBytes := len(splitBytes)

	// Количество добавленных нулевых байт
	countZeroBytes := lenSplitBytes*4 - len(indexBytes)

	for _, a := range splitBytes {
		// Получаем чанки по 3 байта с символами utf-8
		utfBytes := transform4BytesTo3Bytes(a)

		result += string(utfBytes)
	}

	// Возвращаем слайс без нулевых байт
	return result[:len(result)-countZeroBytes]
}

// Удаление символа нулевого байта из конца строки
func trimSymbolZeroByte(str string) string {
	var result string

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == symbolZeroByte[0] {
			continue
		} else {
			result += str[:i+1]
			break
		}
	}

	return result
}

// Разбивка массива байт на чанки по n байт
func splitBytesToByteChunk(bytes []byte, n int) [][]byte {
	var result [][]byte

	j := 0

	chunk := make([]byte, n)

	for i := 0; i < len(bytes); i++ {
		if i%n == 0 && i != 0 {
			result = append(result, chunk)
			chunk = make([]byte, n)
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

// Трансформация буфера из 3 байтов в буфер из 4 байтов
func transform3BytesTo4Bytes(bytes []byte) []byte {
	b1 := bytes[0]
	b2 := bytes[1]
	b3 := bytes[2]

	result := make([]byte, 4)

	// 1 результирующий байт получаем из 6 старших бит 1 байта
	result[0] = extractBites(b1, 2, 7)

	// 2 результирующий байт получаем путем сложения 2 младших бит 1 байта и 4 старших бит 2 байта
	result[1] = extractBites(b1, 0, 1)<<4 | extractBites(b2, 4, 7)

	// 3 результирующий байт получаем путем сложения 4 младших бит 2 байта и 2 старших бит 3 байта
	result[2] = extractBites(b2, 0, 3)<<2 | extractBites(b3, 6, 7)

	// 4 результирующий байт получаем из 6 младших бит 3 байта
	result[3] = extractBites(b3, 0, 5)

	return result
}

// Трансформация буффера из 4 байтов в буффер из 3 байтов
func transform4BytesTo3Bytes(bytes []byte) []byte {
	b1 := bytes[0]
	b2 := bytes[1]
	b3 := bytes[2]
	b4 := bytes[3]

	result := make([]byte, 3)

	// 1 результирующий байт получаем путем сложения 6 младших бит 1 байта и 4-5 бит 2 байта
	result[0] = b1<<2 | extractBites(b2, 4, 5)

	// 2 результирующий байт получаем путем сложения 4 младших бит 2 байта и 2-5 бит 3 байта
	result[1] = b2<<4 | extractBites(b3, 2, 5)

	// 3 результирующий байт получаем путем сложения 2 младших бит 3 байта и 6 младших бит 4 байта
	result[2] = b3<<6 | extractBites(b4, 0, 5)

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
