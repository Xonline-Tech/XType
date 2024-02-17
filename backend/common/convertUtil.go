package common

import "unsafe"

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

// StructToByteSlice 结构体转[]byte
func StructToByteSlice(val any) []byte {
	sizeof := unsafe.Sizeof(val)
	mock := SliceMock{
		uintptr(unsafe.Pointer(&val)),
		int(sizeof),
		int(sizeof),
	}
	return *(*[]byte)(unsafe.Pointer(&mock))
}
