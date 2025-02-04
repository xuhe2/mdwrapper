package core

import (
	"archive/zip"
	"io"
	"os"
	"sync"
)

type ZipArchive struct {
	sync.Mutex

	writer *zip.Writer
}

func NewZipArchive(data *os.File) *ZipArchive {
	return &ZipArchive{
		writer: zip.NewWriter(data),
	}
}

func (za *ZipArchive) Write(src io.Reader, dstPath string) (int64, error) {
	za.Lock()
	defer za.Unlock()

	file, err := za.writer.Create(dstPath)
	if err != nil {
		return 0, err
	}

	return io.Copy(file, src)
}

func (za *ZipArchive) Close() error {
	return za.writer.Close()
}
