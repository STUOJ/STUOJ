package utils_test

import (
	"STUOJ/utils"
	"math/rand/v2"
	"testing"
)

func TestRandKey(t *testing.T) {
	for range 10 {
		l := rand.IntN(256)
		str := utils.GetRandKey(l)
		if len(str) != l {
			t.Errorf("length of string is not equal to length of key")
		} else {
			t.Logf("%s", str)
		}
	}
}
