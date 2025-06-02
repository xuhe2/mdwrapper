package core

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	NormalRefRegex = `\[(?P<label>[^\]]+)\]\((?P<ref>.*?)\)`
	HTMLRefRegex   = `<a href="(?P<ref>[^"]*)"[^>]*>(?P<label>.*?)</a>`

	URLRegex = `(?P<url>https?://[^\s]+)`
)

var RefRegexs = []string{NormalRefRegex, HTMLRefRegex}

type MarkdownFile struct {
	Name    string `json:"name"`    // this is the file name
	Content string `json:"content"` // this is the content of the file
	Path    string `json:"path"`    // this is the absolute path

	offset int // this is the offset of the reader
}

// NewMarkdownFile creates a new MarkdownFile, it is empty
func NewMarkdownFile() *MarkdownFile {
	return &MarkdownFile{}
}

// Open opens a MarkdownFile from a file path
func (f *MarkdownFile) Open(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read all content from the file
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	f.Name = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	f.Path, _ = filepath.Abs(path)
	f.Content = string(content)

	return nil
}

// Read reads from the MarkdownFile, it is a wrapper of io.Reader
func (f *MarkdownFile) Read(p []byte) (n int, err error) {
	if f.offset >= len(f.Content) {
		return 0, io.EOF
	}
	n = copy(p, f.Content[f.offset:])
	f.offset += n
	return n, nil
}

// check if a string is a URL
func isURL(s string) bool {
	return regexp.MustCompile(URLRegex).MatchString(s)
}

// GetFileRefs returns all the file references in the MarkdownFile
//
// no matter it is a normal reference or a HTML reference
func (f *MarkdownFile) GetFileRefs() []string {
	refs := make([]string, 0)
	for _, regex := range RefRegexs {
		re := regexp.MustCompile(regex)
		matches := re.FindAllStringSubmatch(f.Content, -1)
		for _, match := range matches {
			ref := match[re.SubexpIndex("ref")]
			if isURL(ref) {
				continue
			}
			refs = append(refs, ref)
		}
	}
	return refs
}

// GetURLs returns all the URLs in the MarkdownFile
func (f *MarkdownFile) GetURLs() []string {
	refs := make([]string, 0)
	for _, regex := range RefRegexs {
		re := regexp.MustCompile(regex)
		matches := re.FindAllStringSubmatch(f.Content, -1)
		for _, match := range matches {
			ref := match[re.SubexpIndex("ref")]
			if !isURL(ref) {
				continue
			}
			refs = append(refs, ref)
		}
	}
	return refs
}

func (f *MarkdownFile) Replace(old, new string) {
	f.Content = strings.ReplaceAll(f.Content, old, new)
}
