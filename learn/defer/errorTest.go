package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// 自定义错误类型
type FileError struct {
	Op   string
	Path string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("文件操作错误：%s %s：%v", e.Op, e.Path, e.Err)
}

func (e *FileError) Unwrap() error {
	return e.Err
}

// 模拟文件读取函数
func ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return &FileError{
			Op:   "打开",
			Path: path,
			Err:  err,
		}
	}
	defer file.Close()

	buf := make([]byte, 1024)
	_, err = file.Read(buf)
	if err != nil {
		if err == io.EOF {
			return nil // 正常结束
		}
		return &FileError{
			Op:   "读取",
			Path: path,
			Err:  err,
		}
	}
	return nil
}

func main() {
	err := ReadFile("不存在的文件.txt")
	if err != nil {
		// 使用errors.As提取错误类型
		var fileErr *FileError
		if errors.As(err, &fileErr) {
			fmt.Printf("操作：%s，路径：%s，错误：%v\n", fileErr.Op, fileErr.Path, fileErr.Err)
		} else {
			fmt.Println("未知错误：", err)
		}
	} else {
		fmt.Println("文件读取成功")
	}
}
