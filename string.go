/*
  string with atomic store and load operations

  Copyright 2021 Nicolas S. Dade
*/

package atomicstring

import (
	"sync/atomic"
	"unsafe"
)

type String struct {
	ptr *string
}

func (str *String) Store(s string) {
	var ptr unsafe.Pointer
	if s != "" {
		ptr = unsafe.Pointer(&s)
	}
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&str.ptr)), ptr)
}

func (str *String) Load() string {
	ptr := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&str.ptr)))
	if ptr != nil {
		return *(*string)(ptr)
	}
	return ""
}

func (str *String) String() string {
	return str.Load()
}
