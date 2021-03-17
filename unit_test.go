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

	s.Store("abc")
	if s.Load() != "abc" {
		t.Error(`Load(Store("abc")) != "abc"`)
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
