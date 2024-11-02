package base64

import "testing"

type TestData struct {
	input    string
	expected string
}

func TestEncode(t *testing.T) {
	testData := []TestData{
		{
			input:    "hello",
			expected: "aGVsbG8=",
		},
		{
			input:    "world",
			expected: "d29ybGQ=",
		},
		{
			input:    "hello, world!",
			expected: "aGVsbG8sIHdvcmxkIQ==",
		},
		{
			input:    "one",
			expected: "b25l",
		},
		{
			input:    "привет",
			expected: "0L/RgNC40LLQtdGC",
		},
	}

	for _, data := range testData {
		result := Encode(data.input)

		if result != data.expected {
			t.Errorf("input: %s, expected %s, got %s", data.input, data.expected, result)
		}
	}
}

func TestDecode(t *testing.T) {
	testData := []TestData{
		{
			input:    "aGVsbG8=",
			expected: "hello",
		},
		{
			input:    "d29ybGQ=",
			expected: "world",
		},
		{
			input:    "aGVsbG8sIHdvcmxkIQ==",
			expected: "hello, world!",
		},
		{
			input:    "b25l",
			expected: "one",
		},
		{
			input:    "0L/RgNC40LLQtdGC",
			expected: "привет",
		},
	}

	for _, data := range testData {
		result := Decode(data.input)

		if result != data.expected {
			t.Errorf("input: %s, expected %s, got %s", data.input, data.expected, result)
		}
	}
}
