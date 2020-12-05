package expenses

import (
	"io/ioutil"
	"testing"
)

func BenchmarkCalculate(b *testing.B) {
	data, _ := ioutil.ReadFile("./input.txt")

	for i := 0; i < b.N; i++ {
		Calculate(data)
	}
}
