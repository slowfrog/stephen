package model

import (
	"fmt"
	"testing"
)

func TestAlignmentNames(t *testing.T) {
	if HORIZONTAL.Name() != "horizontal" {
		t.Errorf("Wrong name for HORIZONTAL, was %s\n", HORIZONTAL.Name())
	}
	if VERTICAL.Name() != "vertical" {
		t.Errorf("Wrong name for VERTICAL, was %s\n", VERTICAL.Name())
	}
}

func TestSausageToString(t *testing.T) {
	s := CreateSausage(HORIZONTAL, 5, 4)
	s.Cook(1).Cook(3).Cook(3).Cook(3)
	const exp string = "(5,4-horizontal-[.x][.#])"
	str := fmt.Sprintf("%v", s)

	if str != exp {
		t.Errorf("Wrong ToString, expected %s but was %s\n", exp, str)
	}
}
