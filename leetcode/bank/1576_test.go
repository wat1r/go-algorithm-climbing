package test

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	modifyString("dd")
}

func modifyString(s string) string {
	res := []byte(s)
	n := len(s)
	for i, c := range res {
		if c == '?' {
			for b := byte('a'); b <= 'c'; b++ {
				if !(i > 0 && res[i-1] == b || i < n-1 && res[i+1] == b) {
					res[i] = b
					break
				}
			}
		}
	}
	return string(res)
}
