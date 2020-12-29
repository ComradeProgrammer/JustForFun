package helloworld

import "testing"

func TestIsEven(t *testing.T) {
	for i := 0; i < 10; i++ {
		if tmp := IsEven(i); tmp != (i%2 == 0) {
			t.Errorf("Iseven(%d)=%v,exected %v", i, tmp, i%2 == 0)
		}
	}
}

func BenchmarkIsEven(t *testing.B) {
	for i := 0; i < 10; i++ {
		IsEven(i)
	}
}
