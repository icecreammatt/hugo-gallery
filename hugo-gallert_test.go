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
		ImagePath:        "test_path/test_file0.jpg",
		Date:             "2006-01-02",
		NextImage:        "test_path/test_file1.jpg",
		PreviousImage:    "",
		NextPostPath:     "test/test_file1",
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
	postList, err := ioutil.ReadDir("test_data/file_list")
	if err != nil {
		t.Error("Expected files in test_data/file_list", err)
	}

	firstPath := buildPathFromFileInfo(postList[0], "test_data/file_list/", true, "")
	if firstPath != "test_data/file_list/image1" {
		t.Error("Expected", "test_data/file_list/image1", "got", firstPath)
	}

	firstPath = buildPathFromFileInfo(postList[0], "test_data/file_list/", false, "")
	if firstPath != "test_data/file_list/image1.jpg" {
		t.Error("Expected", "test_data/file_list/image1.jpg", "got", firstPath)
	}

	firstPath = buildPathFromFileInfo(postList[0], "test_data/file_list/", false, "s3.amazon.com")
	if firstPath != "s3.amazon.com/test_data/file_list/image1.jpg" {
		t.Error("Expected", "s3.amazon.com/test_data/file_list/image1.jpg", "got", firstPath)
	}

}

func TestGetPreviousAndNextPost(t *testing.T) {
	postList, err := ioutil.ReadDir("test_data/file_list")
	if err != nil {
		t.Error("Expected files in test_data/file_list", err)
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
