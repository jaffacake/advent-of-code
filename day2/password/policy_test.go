package password

import (
	"io/ioutil"
	"testing"
)

type testCase struct {
	p policy
	expectedResult bool
}

func BenchmarkValidatePasswords(b *testing.B) {
	data, _ := ioutil.ReadFile("./input.txt")

	for i := 0; i < b.N; i++ {
		ValidatePasswords(data)
	}
}

func TestSecondPassword(t *testing.T) {
	testCases := make(map[int]testCase)

	testCases[0] = testCase{
		p:              policy{
			int1: 1,
			int2: 3,
			letter: "a",
			password: []byte("abcde"),
		},
		expectedResult: true,
	}

	testCases[1] = testCase{
		p:              policy{
			int1: 1,
			int2: 3,
			letter: "b",
			password: []byte("cdefg"),
		},
		expectedResult: false,
	}

	testCases[2] = testCase{
		p:              policy{
			int1: 2,
			int2: 9,
			letter: "c",
			password: []byte("ccccccccc"),
		},
		expectedResult: false,
	}

	for _, test := range testCases {
		if test.p.testSecondPolicy() != test.expectedResult {
			t.Fail()
		}
	}
}

