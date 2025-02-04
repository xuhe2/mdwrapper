package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Wrapper struct {
	filePathSet *Set[string] // a set for file paths
	archive     *ZipArchive
}

func NewWrapper() *Wrapper {
	return &Wrapper{
		filePathSet: NewSet[string](),
		archive:     nil,
	}
}

func (w *Wrapper) WithArchive(archive *ZipArchive) *Wrapper {
	w.archive = archive
	return w
}

func (w *Wrapper) Wrap(md *MarkdownFile) error {
	fmt.Printf("wrapping file: %s\n", md.Name)

	refs := md.GetFileRefs()
	for _, path := range refs {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		// check if the file has been wrapped
		if w.filePathSet.Contains(absPath) {
			continue
		}
		// wrap the file
		// add the file path to the set
		w.filePathSet.Add(absPath)
		dstPath := filepath.Join("./files/", filepath.Base(absPath))
		if err := w.wrapFile(absPath, dstPath); err != nil {
			return err
		}
		md.Replace(path, dstPath)
	}

	return w.wrapFileFromReader(md, filepath.Join("./", md.Name))
}

func (w *Wrapper) wrapFile(srcPath, dstPath string) error {
	// if archive is nil, return error
	if w.archive == nil {
		return fmt.Errorf("no archive provided")
	}
	// wrap file
	fmt.Printf("wrapping file from %s -> %s\n", srcPath, dstPath)
	// open src file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	w.wrapFileFromReader(srcFile, dstPath)

	return nil
}

func (w *Wrapper) wrapFileFromReader(srcReader io.Reader, dstPath string) error {
	// if archive is nil, return error
	if w.archive == nil {
		return fmt.Errorf("no archive provided")
	}

	w.archive.Write(srcReader, dstPath)

	return nil
}

func (w *Wrapper) Close() {
	if w.archive == nil {
		return
	}
	if err := w.archive.Close(); err != nil {
		fmt.Println(err)
	}
}
