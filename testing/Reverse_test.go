package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)

		require.Equal(t, orig, doubleRev)
		//if utf8.ValidString(orig) && !utf8.ValidString(rev) {
		//	t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		//}
		require.Condition(t, func() (success bool) {
			return !utf8.ValidString(orig) || utf8.ValidString(rev)
		}, fmt.Sprintf("Reverse produced invalid UTF-8 string, orig: %s, rev: %s", orig, rev))
	})
}
