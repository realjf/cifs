package ahocorasick

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewAhoCorasickDoubleArrayTrie(t *testing.T) {
	trie := NewAhoCorasickDoubleArrayTrie()
	files, err := ioutil.ReadFile("./dictionary.txt")
	if err != nil {
		t.Fatal(err)
	}
	fileData := strings.Split(string(files), "\n")
	gMap := gmap.NewTreeMap(gutil.ComparatorString, true)
	for _, d := range fileData {
		gMap.Set(d, d)
	}
	t.Log(gMap.Values())
	trie.Build(*gMap)
	t.Fatal(trie)
}

