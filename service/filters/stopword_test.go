package filters

import "testing"

func TestNewStopword(t *testing.T) {
	stopword := NewStopWord()
	stopword.LoadTable("./stopword_table")
	t.Fatalf("%d", stopword.GetLength())
}
