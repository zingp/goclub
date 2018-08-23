package test

import(
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Fatal("Add is not right")
		return
	}
	t.Logf("Add is right")
}