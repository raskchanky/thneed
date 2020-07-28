package repl

import (
	"testing"
)

func TestRead(t *testing.T) {
	x := "lol"
	if r := Read(x); r != x {
		t.Fatalf("expected %s, got %s", x, r)
	}
}

func TestEval(t *testing.T) {
	x := "lol"
	if r := Eval(x); r != x {
		t.Fatalf("expected %s, got %s", x, r)
	}
}

func TestPrint(t *testing.T) {
	x := "lol"
	if r := Print(x); r != x {
		t.Fatalf("expected %s, got %s", x, r)
	}
}

func TestRep(t *testing.T) {
	x := "lol"
	if r := Rep(x); r != x {
		t.Fatalf("expected %s, got %s", x, r)
	}
}
