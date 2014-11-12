package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := os.Args
	if len(path) < 3 {
		fmt.Printf("Usage: %s <Path> <Section>", path[0])
	}
	fileInfo, err := ioutil.ReadDir(path[1])
	if err != nil {
		panic(err)
	}

    contentPath := "content/" + path[2]
    os.Create(contentPath)
	for index, file := range fileInfo {
		fmt.Printf("%d %s\n", index, file.Name())
        ioutil.WriteFile(contentPath + "/" + file.Name(), []byte("test"), 0644)
	}
}
