package test

import (
	"fmt"
	"gotool/rand"
	"testing"
)

func TestRand(t *testing.T) {
	uuid := rand.NewUUID()
	println(uuid.ToString())
	parserString, err := rand.ParserString(uuid.ToString())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(parserString.ToString())
}

