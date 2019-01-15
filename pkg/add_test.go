package pkg

import (
	"testing"
)

func Test_Add_1(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add fail")
	}
}

func Test_Add_2(t *testing.T) {
	if Add(-1, 5) != 4 {
		t.Error("Add fail")
	}
}
