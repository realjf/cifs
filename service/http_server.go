package main

import (
	"cifs/service/config"
	"cifs/service/db"
	"cifs/service/dictionary"
	"cifs/service/filters"
	"cifs/service/utils"
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"github.com/huichen/sego"
)


const (
	port = ":40001"
)


var (
	Config *config.Config
	pprof_address = "0.0.0.0:6060"

	StopWordTable *filters.StopWord
	SensitiveTable *filters.SensitiveWord
	SimplifiedChineseDict *dictionary.SimplifiedChineseDictionary
	SensitiveWordDict *dictionary.SensitiveWordDictionary

	Segmenter *sego.Segmenter
)

func init() {
	// config、db
	Config = config.NewConfig().LoadConfig("config/config.json")
	db.NewMysql(Config).Init()

	// 过滤多余空格


	// 加载停用词
	StopWordTable = filters.NewStopWord()
	err := StopWordTable.LoadTable("./filters/stopword_table")
	if err != nil {
		log.Printf("failed to load stop word table: %v", err)
	}

	// 加载敏感词
	SensitiveTable = filters.NewSensitiveWord()
	err = SensitiveTable.LoadTable("./filters/sensitiveword_table")
	if err != nil {
		log.Printf("failed to load sensitive word table: %v", err)
	}

	// 加载自定义的敏感词字典
	SensitiveWordDict = dictionary.NewSensitiveWordDictionary()
	err = SensitiveWordDict.LoadDir("../data/dictionary/sensitiveword/", " ")
	if err != nil {
		log.Println(err)
	}
	//SensitiveWordDict = dictionary.NewSensitiveWordDictionary()
	//err = SensitiveWordDict.LoadWith("../data/dictionary/sensitiveword/keywords", " ")
	//if err != nil {
	//	log.Println(err)
	//}

	// 加载繁体转简体的字典
	SimplifiedChineseDict = dictionary.NewSimplifiedChineseDictionary()
	file1 := "../data/dictionary/tc/fanti2jianti.txt"
	err = SimplifiedChineseDict.LoadWith(file1, "\t")
	if err != nil {
		log.Println(err)
	}

	file2 := "../data/dictionary/tc/t2s.txt"
	err = SimplifiedChineseDict.LoadWith(file2, "=")
	if err != nil {
		log.Println(err)
	}

	// 分词词典
	Segmenter.LoadDictionary("../data/dictionary/dictionary.txt")

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		if err := http.ListenAndServe(pprof_address, nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	srv := http.Server{
		Addr: port,
	}
	http.HandleFunc("/filter", filter)
	log.Fatal(srv.ListenAndServe())
}

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Result map[string]interface{} `json:"result"`
}

func (resp *Response) ToJson() []byte {
	res, err := json.Marshal(*resp)
	if err != nil {
		log.Println("json marshal error: ", err)
		return nil
	}

	return res
}

func filter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	content := r.Form.Get("content")
	resp := &Response{}
	resp.Result = map[string]interface{}{
		"original_content": content,
	}
	w.WriteHeader(http.StatusOK)
	if content == "" {
		resp.Code = 101
		resp.Message = "缺少参数"
		w.Write(resp.ToJson())
		return
	}

	// 全角转半角
	content = utils.D2SConvertString(content)

	// 繁体转换简体
	content = SimplifiedChineseDict.TransformString(content)

	// 过滤停用词
	res := StopWordTable.Filter(content)
	log.Printf("filtered_content: %s", res)

	// 过滤敏感词
	res = SensitiveTable.Filter(content)
	log.Printf("filtered_content: %s", res)

	// 自定义敏感词过滤
	content = SensitiveWordDict.Filter(content)

	// 中文分词
	segs := Segmenter.Segment([]byte(content))

	resp.Code = 100
	resp.Message = "过滤成功"
	resp.Result["filtered_content"] = res
	resp.Result["segs"] = segs
	w.Write(resp.ToJson())
}