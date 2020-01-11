package filters

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

// 停用词过滤

type StopWord struct {
	table []string
	sum   int
	filePath string
}

func NewStopWord() *StopWord {
	return &StopWord{
		table: []string{},
	}
}

func (s *StopWord) LoadTable(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	s.filePath = path

	// 加载默认
	s.table = s.LoadDefaultTable()
	table := strings.Split(string(data), "\n")
	s.table = append(s.table, table...)
	s.sum = len(s.table)
	return nil
}

func (s *StopWord) GetLength() int {
	return s.sum
}

func (s *StopWord) GetTable() []string {
	return s.table
}

func (s *StopWord) Add(str string) error {
	if str == "" {
		return errors.New("attempt to add empty string")
	}
	s.table = append(s.table, str)
	return nil
}

func (s *StopWord) Remove(str string) error {
	if str == "" {
		return errors.New("attempt to add empty string")
	}

	return nil
}

func (s *StopWord) Filter(str string) string {
	old := str
	for _, word := range s.table {
		str = strings.Replace(str, word, "", -1)
	}
	log.Printf("content '%s' convert to '%s' ", old, str)
	return str
}

func (s *StopWord) LoadDefaultTable() []string {
	return []string{
		",",
		".",
		"。",
		"，",
		"...",
		"……",
		"~",
		"'",
		"!",
		"！",
		"?",
		"？",
		"：",
		":",
		"——",
		"\"",
		"”",
		"“",
		"@",
		"+",
		"-",
		"_",
		"(",
		")",
		"（",
		"）",
		"|",
		"#",
		"|",
		"%",
		"$",
		"￥",
		"*",
		"=",
		"&",
		"《",
		"》",
		"[",
		"]",
		"{",
		"}",
		"【",
		"】",
		"<",
		">",
		";",
		"；",
		"`",
		"^",
		"/",
		"\\",
	}
}
