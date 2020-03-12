package simplemath

import "testing"

func TestingSqrt1(t *testing.T) {
    v := Sqrt(16)
    if v != 4 {
        t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
    }
}
