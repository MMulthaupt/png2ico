package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/biessek/golang-ico"
)

func main() {
	inFilePath, outFilePath := parseFlags()
	png := mustDecodePng(inFilePath)
	mustWriteToIco(outFilePath, png)
}

func mustDecodePng(pngFilePath string) image.Image {
	inFile := mustOpen(pngFilePath, false)
	defer inFile.Close()
	png, err := png.Decode(inFile)
	if err != nil {
		fatalf("Could not decode PNG: %v", err)
	}
	return png
}

func mustWriteToIco(outFilePath string, img image.Image) {
	outFile := mustOpen(outFilePath, true)
	defer outFile.Close()
	err := ico.Encode(outFile, img)
	if err != nil {
		fatalf("Could not encode ICO: %v", err)
	}
}

func parseFlags() (inFilePath, outFilePath string) {
	inFilePathPtr := flag.String("i", "", "Input PNG file path.")
	outFilePathPtr := flag.String("o", "", "Output ICO file path.")
	flag.Parse()
	inFilePath, outFilePath = *inFilePathPtr, *outFilePathPtr
	if len(inFilePath) == 0 {
		fatalf("Input file path missing.")
	}
	if len(outFilePath) == 0 {
		fatalf("Output file path missing.")
	}
	return
}

func mustOpen(filePath string, write bool) (file *os.File) {
	var err error
	if write {
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			fatalf("Could not open file \"%s\" for writing: %v", filePath, err)
		}
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			fatalf("Could not open file \"%s\" for reading: %v", filePath, err)
		}
	}
	return
}

func fatalf(formatMessage string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf(formatMessage, args...) + "\n")
	os.Exit(1)
}
