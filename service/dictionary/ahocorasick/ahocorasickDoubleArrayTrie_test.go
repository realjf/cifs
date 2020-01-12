package ahocorasick

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
	"strings"
	"testing"
)

func TestNewAhoCorasickDoubleArrayTrie(t *testing.T) {
	trie := NewAhoCorasickDoubleArrayTrie()
	//files, err := ioutil.ReadFile("./dictionary.txt")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fileData := strings.Split(string(files), "\n")
	files := "何炅\n击破\n耳机\n我怕热\n机票\n日期\n金额\n无人接\n企鹅\n破发\n紧迫感\n呢肉\n平稳\n热捧\n体积\n为\n儿童\n玩儿"
	fileData := strings.Split(string(files), "\n")
	gMap := gmap.NewTreeMap(gutil.ComparatorString, true)
	for _, d := range fileData {
		gMap.Set(d, d)
	}
	//t.Log(gMap.Values())
	trie.Build(*gMap)

	t.Log(trie.ExactMatchSearch("击破"))
	hits := trie.ParseText("击破")
	for _, hit := range hits.FrontAll() {
		h := hit.(*Hit)
		t.Log(h.ToString())
		t.Logf("[%d:%d]=%s\n", h.begin, h.end, h.value)
	}
	t.Fatal(1)
}

