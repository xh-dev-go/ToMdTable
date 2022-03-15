package main

import "testing"

func TestValidation(t *testing.T){
	x := [][]string{{"sad", "xxx"}, {"ssssss", "s"}, {"ddd", "sd"}}
	err := validation(x)
	if err != nil {
		t.Error("Unexpected resutl")
	}

	x = [][]string{{"sad", "xxx"}, {"ssssss", "s",""}, {"ddd", "sd"}}
	err = validation(x)
	if err != nil {
		t.Error("expect result")
	} else {
		if err != ColumnCountNotMatch {
			t.Error("expect column count not match")
		}
	}
}

func TestSplitting(t *testing.T) {
	s := "a\tb\nxx\tbb\n"
	msgs, _ := splitting(s)
	if len(msgs) != 2 {
		t.Error("row count not correct")
	} else {
		if len(msgs[0]) != 2 {
			t.Error("first row value count not correct")
		}
		if len(msgs[1]) != 2 {
			t.Error("second row value count not correct")
		}
	}

	s = ""
	_, err := splitting(s)
	if err != MsgIsEmpty {
		t.Error("Not correct msg")
	}

}

func TestGetColumn(t *testing.T) {
	x := [][]string{{"sad", "xxx"}, {"ssssss", "s"}, {"ddd", "sd"}}
	result, _ := getColumnWidth(x)
	if len(result) != 2 {
		t.Error("column not correct")
	}

	if result[0] != 6 {
		t.Error("first column width incorrect")
	}

	if result[1] != 3 {
		t.Error("second column width incorrect")
	}

	x = [][]string{{"sad", "xxx"}, {"ssssss", "s", ""}, {"ddd", "sd"}}
	result, err := getColumnWidth(x)
	if err == nil {
		t.Error("show prompt error")
	}

}
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
