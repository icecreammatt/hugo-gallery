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
previous_image: {{.PreviousImage}}
next_image: {{.NextImage}}
next_post_path: {{.NextPostPath}}
previous_post_path: {{.PreviousPostPath}}
---
`

type GalleryItem struct {
	Title            string
	Date             string
	ImagePath        string
	Weight           string
	NextImage        string
	PreviousImage    string
	NextPostPath     string
	PreviousPostPath string
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
	section := os.Args[2] + "/"
	title := os.Args[3]
	contentPath := "content/" + section

	src, err := os.Stat(contentPath)
	if err != nil || !src.IsDir() {
		err = os.Mkdir(contentPath, 0755)
		check(err)
	}

	postList, err := ioutil.ReadDir(sourcePath)
	check(err)

	for index, file := range postList {
		previousImage, nextImage := getPreviousAndNextPost(index, postList)
		generatePost(index, file, staticRoot, contentPath, title, previousImage, nextImage, section)
	}
}

func getPreviousAndNextPost(index int, postList []os.FileInfo) (previous os.FileInfo, next os.FileInfo) {
	if index+1 < len(postList) {
		next = postList[index+1]
	}
	if index >= 1 {
		previous = postList[index-1]
	}
	return
}

func stripExtension(baseUri string) (fileName string) {
	extensionIndex := strings.Index(baseUri, ".")
	fileName = baseUri[:extensionIndex]
	return
}

func buildPathFromFileInfo(imageFile os.FileInfo, sourcePath string, excludeExtension bool) (imagePath string) {
	if imageFile != nil {
		fileName := imageFile.Name()
		if excludeExtension {
			fileName = stripExtension(imageFile.Name())
		}
		imagePath = sourcePath + fileName
	}
	return
}

func generatePost(index int, file os.FileInfo, sourcePath string, contentPath string, title string, previousImage os.FileInfo, nextImage os.FileInfo, section string) {
	nextImagePath := buildPathFromFileInfo(nextImage, sourcePath, false)
	previousImagePath := buildPathFromFileInfo(previousImage, sourcePath, false)
	nextPostPath := buildPathFromFileInfo(nextImage, section, true)
	previousPostPath := buildPathFromFileInfo(previousImage, section, true)

	galleryItem := GalleryItem{
		Title:            title,
		ImagePath:        sourcePath + file.Name(),
		Date:             time.Now().Format("2006-01-02"),
		Weight:           strconv.Itoa(index),
		NextImage:        nextImagePath,
		PreviousImage:    previousImagePath,
		NextPostPath:     nextPostPath,
		PreviousPostPath: previousPostPath,
	}

	var buffer bytes.Buffer
	generateTemplate(galleryItem, &buffer)

	filePath := contentPath + stripExtension(file.Name()) + ".md"
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
