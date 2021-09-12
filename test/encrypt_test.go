package test

import (
	"go-toolkit/encrypt"
	"testing"
)

func TestMD5(t *testing.T) {
	println(encrypt.MD5("123456"))
}
