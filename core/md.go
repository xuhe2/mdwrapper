package core

import (
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	NormalRefRegex = `\[(?P<label>[^\]]+)\]\((?P<ref>.*)\)`
	HTMLRefRegex   = `<a href="(?P<ref>[^"]*)"[^>]*>(?P<label>.*?)</a>`
)

var RefRegexs = []string{NormalRefRegex, HTMLRefRegex}

type MarkdownFile struct {
	Name    string `json:"name"`    // this is the file name
	Content string `json:"content"` // this is the content of the file
	Path    string `json:"path"`    // this is the absolute path

	offset int // this is the offset of the reader
}

func NewMarkdownFile() *MarkdownFile {
	return &MarkdownFile{}
}

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

func (f *MarkdownFile) Read(p []byte) (n int, err error) {
	if f.offset >= len(f.Content) {
		return 0, io.EOF
	}
	n = copy(p, f.Content[f.offset:])
	f.offset += n
	return n, nil
}

func isURL(s string) bool {
	if _, err := url.ParseRequestURI(s); err == nil {
		return true
	}
	return false
}

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
