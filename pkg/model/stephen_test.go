package model

import (
	"fmt"
	"testing"
)

func TestStephen(t *testing.T) {
	s := Stephen{Pos{5, 4}, UP}
	exp := "{X:5 Y:4 Dir:up}"
	act := fmt.Sprintf("%v", s)
	if act != exp {
		t.Errorf("Wrong stephen, expected\n%s, got\n%s", exp, act)
	}
}
