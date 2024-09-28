package file

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	Year  = "2006"
	Month = "2006-01"
	Day   = "2006-01-02"
	Hour  = "2006-01-02 15"
	Min   = "2006-01-02 15:04"
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

func FilePath(filename string, path string, kind string) string {
	return filepath.Join(path, time.Now().Format(kind), filename)
}

func Ext(filename string) (name, ext string) {
	ext = filepath.Ext(filename)
	return ext, strings.TrimSuffix(filename, ext)
}
