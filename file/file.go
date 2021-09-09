package file

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
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
