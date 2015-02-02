package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"syscall"
	"text/template"
	"time"
)

var postTemplate string = `---
title: {{.Title}}
date: "{{.Date}}"
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
	NextImage        string
	PreviousImage    string
	NextPostPath     string
	PreviousPostPath string
}

func check(e error) int {
	result := 0
	if e != nil {
		result = 1
		defer func() {
			panic(e)
		}()
	}
	return result
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <Source Path> <Destination Section> <Title> [BaseUrl]\n", os.Args[0])
		syscall.Exit(1)
	}

	sourcePath := os.Args[1]
	staticRoot := strings.Replace(os.Args[1], "static/", "", 1) + "/"
	section := os.Args[2] + "/"
	title := os.Args[3]
	baseUrl := os.Args[4]
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
		generatePost(index, file, staticRoot, contentPath, title, previousImage, nextImage, section, baseUrl)
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
	extension := path.Ext(baseUri)
	fileName = baseUri[0 : len(baseUri)-len(extension)]
	return
}

func buildPathFromFileInfo(imageFile os.FileInfo, sourcePath string, excludeExtension bool, baseUrl string) (imagePath string) {
	if imageFile != nil {
		fileName := imageFile.Name()
		if excludeExtension {
			fileName = stripExtension(imageFile.Name())
		}
		if baseUrl != "" && !excludeExtension {
			imagePath = baseUrl + "/" + sourcePath + fileName
		} else {
			imagePath = sourcePath + fileName
		}
	}
	return
}

func generatePost(index int, file os.FileInfo, sourcePath string, contentPath string, title string, previousImage os.FileInfo, nextImage os.FileInfo, section string, baseUrl string) {
	nextImagePath := buildPathFromFileInfo(nextImage, sourcePath, false, baseUrl)
	previousImagePath := buildPathFromFileInfo(previousImage, sourcePath, false, baseUrl)
	nextPostPath := buildPathFromFileInfo(nextImage, section, true, baseUrl)
	previousPostPath := buildPathFromFileInfo(previousImage, section, true, baseUrl)
	currentImagePath := ""
	if baseUrl != "" {
		currentImagePath = baseUrl + "/" + sourcePath + file.Name()
	} else {
		currentImagePath = sourcePath + file.Name()
	}

	galleryItem := GalleryItem{
		Title:            title,
		ImagePath:        currentImagePath,
		Date:             time.Now().Format("2006-01-02"),
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
