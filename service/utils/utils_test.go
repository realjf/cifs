package utils

import "testing"

func TestS2DConvertString(t *testing.T) {
	var s string = "!你好"
	t.Fatal(S2DConvertString(s), s)
}

func TestS2DConvertChar(t *testing.T) {
	var s Char = '!'
	t.Fatal(S2DConvertChar(s), s.ToString())
}


func TestD2SConvertString(t *testing.T) {
	var s string = "你好！"
	t.Fatal(D2SConvertString(s), s)
}

func TestD2SConvertChar(t *testing.T) {
	var s Char = '!'
	t.Fatal(D2SConvertChar(s), s.ToString())
}
