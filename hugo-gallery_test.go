package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestGeneratePost(t *testing.T) {

}

func TestGenerateTemplate(t *testing.T) {
	testItem := GalleryItem{
		Title:            "test_title",
		ImagePath:        "sample-site/static/images/image1.jpg",
		Date:             "2006-01-02",
		NextImage:        "sample-site/static/images/image2.jpg",
		PreviousImage:    "",
		NextPostPath:     "test/image1",
		PreviousPostPath: "",
	}
	var buffer bytes.Buffer
	generateTemplate(testItem, &buffer)

}

func TestStripExtension(t *testing.T) {
	testFileName := "mysample.jpg"
	res := stripExtension(testFileName)
	if res != "mysample" {
		t.Error("For", testFileName, "expected", "mysample", "got", res)
	}

	testFileName = "my.sample.jpg"
	res = stripExtension(testFileName)
	if res != "my.sample" {
		t.Error("For", testFileName, "expected", "my.sample", "got", res)
	}
}

func TestBuildPathFromFileInfo(t *testing.T) {
	postList, err := ioutil.ReadDir("sample-site/static/images")
	if err != nil {
		t.Error("Expected files in sample-site/static/images", err)
	}

	firstPath := buildPathFromFileInfo(postList[0], "sample-site/static/images/", true, "")
	if firstPath != "sample-site/static/images/image1" {
		t.Error("Expected", "sample-site/static/images/image1", "got", firstPath)
	}

	firstPath = buildPathFromFileInfo(postList[0], "sample-site/static/images/", false, "")
	if firstPath != "sample-site/static/images/image1.jpg" {
		t.Error("Expected", "sample-site/static/images/image1.jpg", "got", firstPath)
	}

	firstPath = buildPathFromFileInfo(postList[0], "sample-site/static/images/", false, "s3.amazon.com")
	if firstPath != "s3.amazon.com/sample-site/static/images/image1.jpg" {
		t.Error("Expected", "s3.amazon.com/sample-site/static/images/image1.jpg", "got", firstPath)
	}

}

func TestGetPreviousAndNextPost(t *testing.T) {
	postList, err := ioutil.ReadDir("sample-site/static/images")
	if err != nil {
		t.Error("Expected files in sample-site/static/images", err)
	}

	// Test for start of list
	previousImage, nextImage := getPreviousAndNextPost(0, postList)
	if previousImage != nil {
		t.Error("Expected previousImage", nil, "got", previousImage)
	}
	if nextImage.Name() != "image2.jpg" {
		t.Error("Expected nextImage", "image2.jpg", "got", nextImage.Name())
	}

	// Test for middle of list
	previousImage, nextImage = getPreviousAndNextPost(1, postList)
	if previousImage.Name() != "image1.jpg" {
		t.Error("Expected previousImage", "image1.jpg", "got", previousImage)
	}
	if nextImage.Name() != "image3.jpg" {
		t.Error("Expected nextImage", "image3.jpg", "got", nextImage.Name())
	}

	// Test for end of list
	previousImage, nextImage = getPreviousAndNextPost(3, postList)
	if previousImage.Name() != "image3.jpg" {
		t.Error("Expected previousImage", "image3.jpg", "got", previousImage)
	}
	if nextImage != nil {
		t.Error("Expected nextImage", nil, "got", nextImage)
	}
}

type FakeError struct{}

func (error FakeError) Error() string {
	return "Fake Error"
}

func TestCheck(t *testing.T) {
	var e error
	res := check(e)
	if res != 0 {
		t.Error("Expected 0 for res", res)
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("Recovered check", r)
		}
	}()

	var fakeError FakeError
	res = check(fakeError)
	if res != 0 {
		t.Error("Expected panic for fake error", res)
	}
}
