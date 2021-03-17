package atomicstring

import (
	"testing"
)

func TestString(t *testing.T) {
	var s String

	if s.Load() != "" {
		t.Error(`zero value not ""`)
	}

	s.Store("")
	if s.Load() != "" {
		t.Error(`Load(Store("")) != ""`)
	}
	if s.String() != "" {
		t.Error(`Store("").String() != ""`)
	}

	s.Store("abc")
	if s.Load() != "abc" {
		t.Error(`Load(Store("abc")) != "abc"`)
	}
	if s.String() != "abc" {
		t.Error(`Store("abc").String() != "abc"`)
	}

	s.Store("123")
	if s.Load() != "123" {
		t.Error(`Load(Store("123")) != "123"`)
	}

	s.Store("")
	if s.Load() != "" {
		t.Error(`Load(Store("")) != ""`)
	}
}
