package dictionary

import (
	"cifs/service/dictionary/ahocorasick"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
)

type SensitiveWordDictionary struct {
	CoreDictionary
}

func NewSensitiveWordDictionary() *SensitiveWordDictionary {
	sdic := &SensitiveWordDictionary{
		CoreDictionary{
			files: []string{},
			separator: DefaultSeparator,
			data: gmap.NewTreeMap(gutil.ComparatorString, true),
			dict: ahocorasick.NewAhoCorasickDoubleArrayTrie(),
		},
	}

	return sdic
}

// 过滤
func (swd *SensitiveWordDictionary) Filter(str string) string {
	if str == "" {
		return str
	}

	// 字典过滤

	return str
}

