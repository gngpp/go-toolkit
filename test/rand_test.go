package test

import (
	"fmt"
	"go-toolkit/rand"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := rand.NewUUID()
	println(uuid.ToString())
	parserString, err := rand.ParserString(uuid.ToString())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(parserString.ToString())
}

func TestRand(t *testing.T) {
	randInt := rand.GetInt(5, 100)
	fmt.Println(randInt)
	randomAscii := rand.RandomAscii(10)
	fmt.Println(randomAscii)
	getString1 := rand.GetBytes(10, rand.ALL)
	fmt.Println(string(getString1))
	getString2 := rand.GetBytes(10, rand.NUM)
	fmt.Println(string(getString2))
	getString3 := rand.GetBytes(10, rand.LOWER)
	fmt.Println(string(getString3))
	getString4 := rand.GetBytes(10, rand.UPPER)
	fmt.Println(string(getString4))
}
