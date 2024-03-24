package filex

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Create 创建文件
func Create(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func RemoveFile(filename string) error {
	return os.Remove(filename)
}

func RemovePath(path string) error {
	return os.RemoveAll(path)
}
