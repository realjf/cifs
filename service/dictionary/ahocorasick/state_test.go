package ahocorasick

import (
	. "asifs/service/utils"
	"testing"
)

func TestNewState(t *testing.T) {
	state := NewState()
	s := "我净额哦啊牛而紧迫我安全金融IE就温柔而忘记你让我去基恩人迫切而我家人聘请欧威尔"
	ss := String(s).ToCharArray()
	for _, st := range ss {
		state.AddState(st)
	}
	t.Log(state.GetSuccess())
	t.Log(state.GetDepth())
	keys := state.GetTransitions()
	kks := []string{}
	for _, k := range keys {
		kks = append(kks, k.ToString())
	}
	t.Log(kks)
	t.Log()
	t.Fatalf("%v", state)
}
