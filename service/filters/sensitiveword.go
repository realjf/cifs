package filters

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

// 停用词过滤

type SensitiveWord struct {
	table []string
	sum   int
	filePath string
}

func NewSensitiveWord() *SensitiveWord {
	return &SensitiveWord{
		table: []string{},
	}
}

func (s *SensitiveWord) LoadTable(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	s.filePath = path

	table := strings.Split(string(data), "\n")
	s.table = append(s.table, table...)
	s.sum = len(s.table)
	return nil
}

func (s *SensitiveWord) GetLength() int {
	return s.sum
}

func (s *SensitiveWord) GetTable() []string {
	return s.table
}

func (s *SensitiveWord) Add(str string) error {
	if str == "" {
		return errors.New("attempt to add empty string")
	}
	s.table = append(s.table, str)
	return nil
}

func (s *SensitiveWord) Filter(str string) string {
	old := str
	for _, word := range s.table {
		str = strings.Replace(str, word, "", -1)
	}
	log.Printf("content '%s' convert to '%s' ", old, str)
	return str
}

