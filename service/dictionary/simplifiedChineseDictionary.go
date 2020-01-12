package dictionary

import (
	"cifs/service/utils"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
)

type SimplifiedChineseDictionary struct {
	CoreDictionary
}

func NewSimplifiedChineseDictionary() *SimplifiedChineseDictionary {
	sdic := &SimplifiedChineseDictionary{
		CoreDictionary{
			files: []string{},
			separator: DefaultSeparator,
			data: gmap.NewTreeMap(gutil.ComparatorString, true),
		},
	}

	return sdic
}

// 转换繁体字符串
func (scd *SimplifiedChineseDictionary) TransformString(str string) string {
	if str == "" {
		return str
	}

	strChar := utils.String(str).ToCharArray()
	resStr := utils.NewString()
	for _, s := range strChar {
		if o := scd.data.Get(s.ToString()); o != "" {
			resStr.Append(o)
		}else{
			resStr.Append(string(s))
		}
	}

	return resStr.ToString()
}

