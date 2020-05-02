package numtowords

import "testing"

func Test_Transform(t *testing.T) {

	tt := []struct {
		input    int64
		expected string
	}{
		{input: 1234, expected: "one thousand two hundred thirty four"},
		{input: 1100100100300800112, expected: "one quintillion one hundred quadrillion one hundred trillion one hundred billion three hundred million eight hundred thousand one hundred twelve"},
		{input: 100000000000, expected: "one hundred billion"},
		{input: 000000000, expected: "zero"},
		{input: 0, expected: "zero"},
	}

	for _, tc := range tt {
		output := Transform(tc.input)

		if output != tc.expected {
			t.Fatalf("Output should be %s, but got %s", tc.expected, output)
		}
	}
}
