package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
	"text/template"
	"time"
)

var postTemplate string = `---
title: {{.Title}}
date: "{{.Date}}"
weight: {{.Weight}}
image_name: {{.ImagePath}}
---
`

type GalleryItem struct {
	Title     string
	Date      string
	ImagePath string
	Weight    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <Source Path> <Destination Section> <Title>\n", os.Args[0])
		syscall.Exit(1)
	}

	sourcePath := os.Args[1]
	staticRoot := strings.Replace(os.Args[1], "static/", "", 1) + "/"
	section := os.Args[2]
	title := os.Args[3]
	contentPath := "content/" + section + "/"

	src, err := os.Stat(contentPath)
	if err != nil || !src.IsDir() {
		err = os.Mkdir(contentPath, 0755)
		check(err)
	}

	fileInfo, err := ioutil.ReadDir(sourcePath)
	check(err)

	for index, file := range fileInfo {
		generatePost(index, file, staticRoot, contentPath, title)
	}
}

func generatePost(index int, file os.FileInfo, sourcePath string, contentPath string, title string) {
	galleryItem := GalleryItem{
		Title:     title,
		ImagePath: sourcePath + file.Name(),
		Date:      time.Now().Format("2006-01-02"),
		Weight:    strconv.Itoa(index),
	}

	var buffer bytes.Buffer
	generateTemplate(galleryItem, &buffer)

	extensionIndex := strings.Index(file.Name(), ".")
	fileName := file.Name()[:extensionIndex]
	filePath := contentPath + fileName + ".md"
	f, err := os.Create(filePath)
	check(err)
	defer f.Close()
	f.Sync()
	w := bufio.NewWriter(f)
	w.WriteString(buffer.String())
	w.Flush()
}

func generateTemplate(galleryItem GalleryItem, buffer *bytes.Buffer) {

	t := template.New("post template")
	t, _ = t.Parse(postTemplate)
	err := t.Execute(buffer, galleryItem)
	check(err)
}
