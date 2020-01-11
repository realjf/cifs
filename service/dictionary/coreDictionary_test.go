package dictionary

import (
	"testing"
)

func TestNewCoreDictionary(t *testing.T) {
	path := "../../data/dictionary/tc/s2t.txt"

	dict := NewCoreDictionary()
	t.Log(dict.Load(path))
	t.Log(dict.Length())

	// clear
	dict.Clear()
	t.Log(dict.Length())

	path = "../../data/dictionary/tc/t2s.txt"
	t.Log(dict.Load(path))
	t.Log(dict.Length())

	t.Log(dict.dict.ParseText("你好"))

	t.Fatal(1)
}

//func TestReadFiles(t *testing.T) {
//	dirPath := "../../data/dictionary/sensitiveword/"
//	files, err := ioutil.ReadDir(dirPath)
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//
//	if len(files) <= 0 {
//		t.Fatal("目录文件为空")
//	}
//
//	data := map[string]string{}
//
//	for _, f := range files {
//		ff := f.Name()
//		fileData, err := ioutil.ReadFile(dirPath + ff)
//		if err != nil {
//			t.Fatal(err)
//		}
//		t.Log(string(fileData))
//
//		// 文件数据加载
//		sensitiveWords := strings.Split(string(fileData), "\n")
//		for _, kw := range sensitiveWords {
//			value := strings.Split(string(kw), "=")
//			if len(value) >= 2 {
//				data[value[0]] = value[1]
//			}else{
//				data[value[0]] = value[0]
//			}
//		}
//		t.Fatal(data)
//	}
//
//	t.Fatal(1)
//}
