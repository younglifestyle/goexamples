package utils

import (
	"reflect"
	"strconv"
	"unsafe"
)

// StringToSlice string to slice with out data copy
func StringToSlice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

// ToString unsafe 转换, 将 []byte 转换为 string
func ToString(p []byte) string {
	return *(*string)(unsafe.Pointer(&p))
}

// ToBytes unsafe 转换, 将 string 转换为 []byte
func ToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// IntToBool int 类型转换为 bool
func IntToBool(i int) bool {
	return i != 0
}

// SliceInt64ToString []int64 转换为 []string
func SliceInt64ToString(si []int64) (ss []string) {
	ss = make([]string, 0, len(si))
	for k := range si {
		ss = append(ss, strconv.FormatInt(si[k], 10))
	}
	return ss
}

// SliceStringToInt64 []string 转换为 []int64
func SliceStringToInt64(ss []string) (si []int64) {
	si = make([]int64, 0, len(ss))
	var (
		i   int64
		err error
	)
	for k := range ss {
		i, err = strconv.ParseInt(ss[k], 10, 64)
		if err != nil {
			continue
		}
		si = append(si, i)
	}
	return
}

// SliceStringToInt []string 转换为 []int
func SliceStringToInt(ss []string) (si []int) {
	si = make([]int, 0, len(ss))
	var (
		i   int
		err error
	)
	for k := range ss {
		i, err = strconv.Atoi(ss[k])
		if err != nil {
			continue
		}
		si = append(si, i)
	}
	return
}

// MustInt 将string转换为int, 忽略错误
func MustInt(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

// MustInt64 将string转换为int64, 忽略错误
func MustInt64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 64)
	return
}
