package asciiart

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"asciiartweb/internal/asciiartfs"
)

const (
	fileLen = 855
)

// check amount of arguments
func AsciiArt(banner string, fontStr string) (string, error) { // name string, font string

	// Read the content of the file
	argsArr := strings.Split(strings.ReplaceAll(fontStr, "\r", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("../../internal/asciiart/fonts/" + banner + ".txt")
	if err != nil {
		defer readFile.Close()
		return "", errors.New("failed to open file")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		// return "bad code rami -- try again"
		return "", errors.New("file is corrupted")
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	return asciiartfs.PrintBanners(argsArr, arr), nil
}
