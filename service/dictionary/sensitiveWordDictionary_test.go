package dictionary

import (
	"cifs/service/dictionary/ahocorasick"
	"strings"
	"log"
	"testing"
)

func TestNewSensitiveWordDictionary(t *testing.T) {
	var s string = "fuck案件评估IE铺盖卷儿疲软"
	dict := NewSensitiveWordDictionary()
	err := dict.LoadWith("../../data/dictionary/sensitiveword/keywords", " ")
	if err != nil {
		log.Fatalf("%v", err)
	}

	res := dict.dict.ParseText(s)
	if res != nil {
		for _, v := range res.FrontAll() {
			if v == nil {
				continue
			}
			vv := v.(*ahocorasick.Hit)
			s = strings.ReplaceAll(s, vv.Value().(string), "***")
		}
	}
	t.Log(s)
	t.FailNow()
}
