package utils

import "testing"

func TestC2ANumber(t *testing.T) {
	s := "壹佰贰拾叁億肆仟伍佰陆拾柒萬捌仟玖佰零壹元壹角贰分"
	t.Fatal(C2ANumber(s))
}

func TestA2ChinaNumber(t *testing.T) {
	s := "壹佰贰拾叁億肆仟伍佰陆拾柒萬捌仟玖佰零壹元壹角贰分"
	t.Log(C2ANumber(s))
	t.Fatal(A2ChinaNumber(C2ANumber(s)))
}

