package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

const startBufSize = 4096

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	var (
		bufSize              int
		inPath, outPath, sep string
	)

	flag.IntVar(&bufSize, "b", bufio.MaxScanTokenSize, "buffer size of a line")
	flag.StringVar(&inPath, "i", "", "path to an input file")
	flag.StringVar(&outPath, "o", "", "path to an output file")
	flag.StringVar(&sep, "s", " ", "line separator")
	flag.Parse()

	inFile, err := getInFile(inPath)
	if err != nil {
		return err
	}

	defer func() {
		if e := inFile.Close(); e != nil && err == nil {
			err = e
		}
	}()

	outFile, err := getOutFile(outPath)
	if err != nil {
		return err
	}

	defer func() {
		if e := outFile.Close(); e != nil && err == nil {
			err = e
		}
	}()

	_, err = outFile.WriteString(joinLines(inFile, sep, bufSize))
	return err
}

// getInFile returns an input file.
func getInFile(path string) (*os.File, error) {
	if path == "" {
		return os.Stdin, nil
	}
	return os.Open(path)
}

// getOutFile returns an output file.
func getOutFile(path string) (*os.File, error) {
	if path == "" {
		return os.Stdout, nil
	}
	return os.Create(path)
}

// joinLines joins lines of the given input file, using `sep` as the separator.
func joinLines(r io.Reader, sep string, bufSize int) string {
	var elems []string

	scanner := bufio.NewScanner(r)
	buf := make([]byte, startBufSize)
	scanner.Buffer(buf, bufSize)

	for scanner.Scan() {
		elems = append(elems, scanner.Text())
	}

	return strings.Join(elems, sep)
}
