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
	}

	for _, data := range testData {
		result := Encode(data.input)

		if result != data.expected {
			t.Errorf("input: %s, expected %s, got %s", data.input, data.expected, result)
		}
	}
}
