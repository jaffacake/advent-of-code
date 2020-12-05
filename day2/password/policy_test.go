package password

import (
	"io/ioutil"
	"testing"
)

func BenchmarkValidatePasswords(b *testing.B) {
	data, _ := ioutil.ReadFile("./input.txt")

	for i := 0; i < b.N; i++ {
		ValidatePasswords(data)
	}
}
