package core

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestMarkdownFile(t *testing.T) {
	// new a instance
	md := NewMarkdownFile()
	if md == nil {
		t.Error("NewMarkdownFile() failed")
	}
	// open a file
	filePath := "../test/main.md"
	if md = NewMarkdownFile().WithFilePath(filePath); md == nil {
		t.Error("Open markdown file failed")
		return
	}
	if md.Name != strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1] {
		t.Error("Open() failed")
	}
	filePath, _ = filepath.Abs(filePath)
	if md.Path != filePath {
		t.Error("Open() failed")
	}

	fmt.Printf("%+v\n", md.GetFileRefs())
	fmt.Printf("%+v\n", md.GetURLs())
}

func TestIsURL(t *testing.T) {
	url := "https://www.baidu.com"
	filePath := "./main.md"
	linuxAbsFilePath := "/home/main.md"
	windowsAbsFilePath := "D:\\main.md"
	if !isURL(url) {
		t.Errorf("isURL() failed in url %s", url)
	}
	if isURL(filePath) {
		t.Errorf("isURL() failed in filePath %s", filePath)
	}
	if isURL(linuxAbsFilePath) {
		t.Errorf("isURL() failed in linuxAbsFilePath %s", linuxAbsFilePath)
	}
	if isURL(windowsAbsFilePath) {
		t.Errorf("isURL() failed in windowsAbsFilePath %s", windowsAbsFilePath)
	}
}
