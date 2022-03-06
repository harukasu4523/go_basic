package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Exampli of io.SectionREader\n")
	// スタート位置を固定して取り出すio.NewSectionReader(io.ReaderAt型, 開始位置, バイト数(文字数))
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
