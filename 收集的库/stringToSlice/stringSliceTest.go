package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	oneSlice := "21321421421321acxzczxvsadwqdwqfsacadwqfdcasadasd"
	fmt.Println(StringToSlice(oneSlice))

	oneByte := StringToSlice(oneSlice)
	fmt.Println(byteString(oneByte))
}

type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))
}

func byteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// SliceToString slice to string with out data copy
func SliceToString(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// StringToSlice string to slice with out data copy
func StringToSlice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}
