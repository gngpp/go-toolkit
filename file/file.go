package file

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// CurrentDir get the current file dir of the caller
func CurrentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}

// SavaToFile Save the byte array to the file
func SavaToFile(filePath string, bytes []byte) error {
	return ioutil.WriteFile(filePath, bytes, fs.ModePerm)
}

// MakeDir Recursively create new folders based on the incoming folder name
func MakeDir(fileDir string) error {
	return os.MkdirAll(fileDir, fs.ModePerm)
}

// MakeDirByFile create a folder recursively based on the incoming file name
// ./dir/filename  /home/dir/filename
func MakeDirByFile(filePath string) error {
	temp := strings.Split(filePath, "/")
	if len(temp) <= 2 {
		return errors.New("please input complete file name like ./dir/filename or /home/dir/filename")
	}
	dirPath := strings.Join(temp[0:len(temp)-1], "/")
	return MakeDir(dirPath)
}

// Exist determine whether the file exists
func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func Size(path string) int64 {
	stat, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return stat.Size()
}

// InsertSuffix insert a suffix to filepath
func InsertSuffix(src string, suffix string) string {
	ext := path.Ext(src)
	return fmt.Sprintf("%s%s%s", src[:len(src)-len(ext)], suffix, ext)
}

// ReplaceExt replace ext
func ReplaceExt(src string, ext string) string {
	srcExt := path.Ext(src)
	return fmt.Sprintf("%s%s", src[:len(src)-len(srcExt)], ext)
}

// Read file contents
func Read(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	return string(file), err
}
