package errors

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 使用errors.Errorf()函数返回错误的时候不方便定位是那个文件的哪行代码，特设计该功能

var _ error = (*dpipeErr)(nil)

type dpipeErr struct {
	file string
	line int
	err  error
	msg  string
}

// New 实例化一个dpipe错误信息
func New(msg string, err error) error {
	_, file, line, _ := runtime.Caller(1)
	// 文件名称可能比较长，格式化一下
	ff := strings.Split(file, "/")
	if len(ff) > 3 {
		file = path.Join(
			ff[len(ff)-3],
			ff[len(ff)-2],
			ff[len(ff)-1],
		)
	}
	return &dpipeErr{
		file: file,
		line: line,
		err:  err,
		msg:  msg,
	}
}

func Errorf(format string, a ...any) error {
	_, file, line, _ := runtime.Caller(1)
	// 文件名称可能比较长，格式化一下
	ff := strings.Split(file, "/")
	if len(ff) > 3 {
		file = path.Join(
			ff[len(ff)-3],
			ff[len(ff)-2],
			ff[len(ff)-1],
		)
	}
	return &dpipeErr{
		file: file,
		line: line,
		err:  nil,
		msg:  fmt.Sprintf(format, a...),
	}
}

// Error 返回包含错误发生文件名、行号和错误信息的字符串
func (de *dpipeErr) Error() string {
	if de.err == nil {
		return fmt.Sprintf("%s:%d:%s", de.file, de.line, de.msg)
	}
	return fmt.Sprintf("%s:%d:%s::%v", de.file, de.line, de.msg, de.err)
}

// Unwrap 方法用于从包装错误中获取底层错误
func (de *dpipeErr) Unwrap() error {
	return de.err
}

func Is(e, target error) bool {
	return errors.Is(e, target)
}
