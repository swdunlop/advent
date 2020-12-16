package main

import (
	"io/ioutil"
	"os"
)

func main() {
	day := os.Args[1]
	name := os.Args[2]
	err := os.MkdirAll(day+"-"+name, 0700)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(day+"-"+name+"/"+name+".go", []byte("package main\n"), 0600)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(day+"-"+name+"/"+name+"_test.go", []byte("package main\n"), 0600)
	if err != nil {
		panic(err)
	}
}
