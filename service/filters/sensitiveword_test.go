package filters

import "testing"

func TestNewSensitiveWord(t *testing.T) {
	sensitivewd := NewSensitiveWord()
	sensitivewd.LoadTable("./sensitiveword_table")
	t.Fatalf("%d", sensitivewd.GetLength())
}


