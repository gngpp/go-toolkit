package test

import (
	"fmt"
	"go-toolkit/file"
	"go-toolkit/homedir"
	"go-toolkit/logging"
	log2 "log"
	"os"
	"runtime"
	"testing"
)

func TestFile(t *testing.T) {
	println(file.CurrentDir())
	caller, f, line, ok := runtime.Caller(1)
	fmt.Println(caller)
	fmt.Println(f)
	fmt.Println(line)
	fmt.Println(ok)
}

func TestGetHome(t *testing.T) {
	create, err := os.Create("/Users/ant/logs/logging.txt")
	if err != nil {
		log2.Fatal(err)
	}

	err = create.Close()
	if err != nil {
		log2.Fatal(err)
	}
	logging.PrintInfo()
}

func TestDir(t *testing.T) {
	println(homedir.Dir())
	dirs, err := file.GetPathDirs("/Users")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range dirs {
		fmt.Println(v)
	}
}

func TestCreate(t *testing.T) {
	//err := file.Create("/Users/ant/demo","hanbi", "test.txt")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	err := file.DelFile("/Users/ant/demo")
	if err != nil {
		fmt.Println(err.Error())
	}
	println(file.GetCurrentDirectory())
	fmt.Println(file.FormatFileSize(132933635857711104))
}
