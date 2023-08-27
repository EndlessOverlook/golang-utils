package main

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	f, err := os.Create("lkl.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

}

func TestMkdir(t *testing.T) {
	err := os.Mkdir("mk", os.ModePerm)
	if err != nil {
		panic(err)
	}
}
