package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <Path> <Section>\n", os.Args[0])
		syscall.Exit(1)
	}

	contentPath := "content/" + os.Args[2] + "/"
	src, err := os.Stat(contentPath)
	if err != nil || !src.IsDir() {
		err = os.Mkdir(contentPath, 0755)
		check(err)
	}

	fileInfo, err := ioutil.ReadDir(os.Args[1])
	check(err)

	for index, file := range fileInfo {
		fmt.Printf("%d %s\n", index, file.Name())
		filePath := contentPath + file.Name() + ".md"
		fmt.Println(filePath)
		f, err := os.Create(filePath)
		check(err)
		defer f.Close()
		f.Sync()
		w := bufio.NewWriter(f)
		w.WriteString(file.Name())
		w.Flush()
	}
}
