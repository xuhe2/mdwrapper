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
	if err := md.Open(filePath); err != nil {
		t.Error(err)
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
