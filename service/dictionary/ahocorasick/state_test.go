package ahocorasick

import (
	. "cifs/service/utils"
	"fmt"
	"strings"
	"testing"
)

func TestNewState(t *testing.T) {
	state := NewState()
	files := "何炅\n击破\n耳机\n我怕热\n机票\n日期\n金额\n无人接\n企鹅\n破发\n紧迫感\n呢肉\n平稳\n热捧\n体积\n为\n儿童\n玩儿"
	s := strings.Split(string(files), "\n")
	for _, st := range s {
		ss := String(st).ToCharArray()
		root := state
		for _, s1 := range ss  {
			root = root.AddState(s1)
		}
	}
	succ := state.GetSuccess()
	for k, v := range succ {
		str := fmt.Sprintf("%v => ", k.(Char).ToString())
		vv := v.(*State)
		for i :=0; i < vv.GetDepth(); i++ {
			v1 := vv.GetSuccess()
			for kk, v2 := range v1 {
				str += kk.(Char).ToString()
				vv = v2.(*State)
				str += " => "
			}
		}
		t.Log(str)
	}
	t.Log(state.GetSuccess())
	t.Log(state.GetDepth())
	keys := state.GetTransitions()
	kks := []string{}
	for _, k := range keys {
		kks = append(kks, k.ToString())
	}
	t.Log(kks)
	t.Fatalf("%v", state)
}
