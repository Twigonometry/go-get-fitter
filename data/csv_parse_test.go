package data

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestHello(t *testing.T) {
	array := [3]string{"a", "b", "c"}
	want := array[0:3]
    if got := SplitByComma("a, b, c"); !cmp.Equal(got, want) {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}