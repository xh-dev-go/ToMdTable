package main

import "testing"

func TestPad(t *testing.T) {
	var result = pad("abc", 5, " ")
	if result != " abc " {
		t.Error("fail with normal padding")
	}
	result = pad("abc", 6, " ")
	if result != "  abc " {
		t.Error("fail left right padding")
	}
	result = pad("abc", 1, " ")
	if result != "abc" {
		t.Error("fail smaller target padding")
	}
	result = pad("abc", 10, "=")
	if result != "====abc===" {
		t.Error("fail padding specific symbol")
	}
	result = pad("abc", 10, "==")
	if result != "========abc======" {
		t.Error("fail padding specific symbol(multiple rune)")
	}

}
