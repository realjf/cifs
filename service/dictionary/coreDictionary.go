package dictionary

import (
	"cifs/service/dictionary/ahocorasick"
	"errors"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
	"io/ioutil"
	"strings"
)

const (
	DefaultSeparator = "=" // 默认的分隔符
)

// 字典
type CoreDictionary struct {
	files []string
	separator string
	data *gmap.TreeMap
	dict *ahocorasick.AhoCorasickDoubleArrayTrie
}


func NewCoreDictionary() *CoreDictionary {
	return &CoreDictionary{
		files: []string{},
		separator: DefaultSeparator,
		data: gmap.NewTreeMap(gutil.ComparatorString, true),
		dict: ahocorasick.NewAhoCorasickDoubleArrayTrie(),
	}
}

// 默认加载，使用等号（=）分割
func (cd *CoreDictionary) Load(path string) error {
	if path == "" {
		return errors.New("字典路径为空")
	}

	if cd.ContainPath(path) {
		return errors.New("字典已经加载过")
	}

	// 加载文件
	cd.AddPath(path)
	return cd.load(cd.files, DefaultSeparator)
}

func (cd *CoreDictionary) ContainPath(path string) bool {
	for _, f := range cd.files {
		if f == path {
			return true
		}
	}

	return false
}

func (cd *CoreDictionary) AddPath(path string) bool {
	cd.files = append(cd.files, path)
	return true
}

func (cd *CoreDictionary) load(files []string, separator string) error {
	for _, path := range files {
		fileData, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.New(fmt.Sprintf("read file %s error: %v", path, err))
		}
		data := strings.Split(string(fileData), "\n")
		for _, k := range data {
			if k == "" {
				continue
			}
			value := strings.Split(k, separator)
			if len(value) >= 2 {
				cd.data.Set(value[0], value[1])
			}else{
				cd.data.Set(value[0], value[0])
			}
		}
		cd.dict.Build(*cd.data)
		cd.separator = separator
	}

	return nil
}

func (cd *CoreDictionary) LoadDir(dirPath string, separator string) error {
	if dirPath == "" {
		return errors.New("加载目录为空")
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return errors.New(fmt.Sprintf("读取目录下字典错误：%v", err))
	}

	if len(files) <= 0 {
		return errors.New("目录下为空")
	}

	if separator == "" {
		separator = DefaultSeparator
	}

	cd.separator = separator

	for _, f := range files {
		ff := f.Name()
		fileData, err := ioutil.ReadFile(dirPath + ff)
		if err != nil {
			return err
		}

		// 添加文件路径
		if !cd.ContainPath(dirPath + ff) {
			cd.AddPath(dirPath + ff)
		}else{
			continue
		}

		// 文件数据加载
		sensitiveWords := strings.Split(string(fileData), "\n")

		for _, kw := range sensitiveWords {
			value := strings.Split(kw, separator)
			if len(value) >= 2 {
				cd.data.Set(value[0], value[1])
			}else{
				cd.data.Set(value[0], value[0])
			}
		}

		cd.dict.Build(*cd.data)
	}

	return nil
}

func (cd *CoreDictionary) LoadFiles(files []string, separator string) error {
	if len(files) <= 0 {
		return errors.New("没有字典需要加载")
	}

	if separator == "" {
		separator = DefaultSeparator
	}

	cd.separator = separator

	for _, f := range files {
		fileData, err := ioutil.ReadFile(f)
		if err != nil {
			return errors.New("字典读取失败")
		}
		keywords := strings.Split(string(fileData), "\n")
		for _, kw := range keywords {
			value := strings.Split(kw, separator)
			if len(value) >= 2 {
				cd.data.Set(value[0], value[1])
			}else{
				cd.data.Set(value[0], value[0])
			}
		}
		cd.dict.Build(*cd.data)
	}

	return nil
}

// 使用自定义分隔符字典加载方式
func (cd *CoreDictionary) LoadWith(path string, separator string) error {
	if path == "" {
		return errors.New("字典路径不存在")
	}

	if separator == "" {
		// 使用默认的分隔符
		separator = DefaultSeparator
	}

	cd.separator = separator

	if cd.ContainPath(path) {
		return errors.New("字典已经加载过")
	}

	cd.separator = separator
	// 加载文件
	cd.AddPath(path)
	return cd.load(cd.files, separator)
}

// 重新加载
func (cd *CoreDictionary) Reload() error {
	if len(cd.files) <= 0 {
		return errors.New("没有可用的字典路径")
	}
	files := cd.files
	separator := cd.separator
	cd.Clear()
	return cd.load(files, separator)
}

// 清空字典
func (cd *CoreDictionary) Clear() {
	cd.files = []string{}
	cd.data = gmap.NewTreeMap(gutil.ComparatorString, true)
}

func (cd *CoreDictionary) Length() int {
	return cd.data.Size()
}

func (cd *CoreDictionary) GetAll() []interface{} {
	return cd.data.Values()
}

func (cd *CoreDictionary) Get(key string) string {
	if key == "" {
		return ""
	}

	if o := cd.data.Get(key); o != nil {
		return o.(string)
	}else{
		return ""
	}
}

func (cd *CoreDictionary) Add(keyword string) bool {
	if keyword == "" {
		return false
	}

	if o := cd.data.Get(keyword); o != nil {
		return true
	}else{
		cd.data.Set(keyword, keyword)
	}
	return true
}

func (cd *CoreDictionary) Remove(keyword string) bool {
	if keyword == "" {
		return false
	}

	if cd.data.Get(keyword) != nil {
		cd.data.Remove(keyword)
	}else{
		return false
	}

	return true
}

func (cd *CoreDictionary) write() bool {
	return true
}

