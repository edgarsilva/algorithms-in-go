package reverse_polish_notation

import (
	"testing"
)

func TestAssert(t *testing.T) {
	passes := Assert()

	if !passes {
		t.Errorf("Assert punction failed")
	}
}

func TestEvalMathExp(t *testing.T) {
	exp := "(4*3)+2(3)-5"
	got := EvalMathExp(exp)
	want := "13.00"

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}
