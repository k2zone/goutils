package strlib

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInts format int int32 int64... slice like:n1,n2,n3.
func JoinInts(src interface{}) (s string) {
	var sv reflect.Value
	sv = reflect.ValueOf(src)
	// 判断sv是否是指针
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}
	buf := bfPool.Get().(*bytes.Buffer)
	svLen := sv.Len()
	for i := 0; i < svLen; i++ {
		switch sv.Type().Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			buf.WriteString(strconv.FormatInt(sv.Index(i).Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			buf.WriteString(strconv.FormatInt(int64(sv.Index(i).Uint()), 10))
		}
		buf.WriteString(",")
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s = buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return
}

// SplitInts split string into int8 int int32 int64 uint... slice.
func SplitInts(s string, dist interface{}) (err error) {
	if s == "" {
		return
	}
	dv := reflect.ValueOf(dist)

	if dv.Kind() != reflect.Ptr {
		err = fmt.Errorf("SplitInts: param err, please use slice ptr")
		return
	}
	if dv.Elem().Kind() != reflect.Slice {
		err = fmt.Errorf("SplitInts: param err, please use slice ptr")
		return
	}
	var (
		sliceType = dv.Elem().Type().Elem()
		tar       []reflect.Value
	)

	sArr := strings.Split(s, ",")
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, sliceType.Bits())
		if err != nil {
			return err
		}
		sv := reflect.ValueOf(i)
		if sv.Type().ConvertibleTo(sliceType) {
			tar = append(tar, reflect.ValueOf(i).Convert(sliceType))
		}
	}
	dv.Elem().Set(reflect.Append(dv.Elem(), tar...))
	return
}
