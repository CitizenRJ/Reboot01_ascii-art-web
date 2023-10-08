package asciiart

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"asciiartweb/internal/asciiartfs"
)

const (
	fileLen = 855
)

// check amount of arguments
func AsciiArt(banner string, fontStr string) string { // name string, font string

	// Validate the input string
	err := asciiartfs.IsValid(fontStr)
	if err != nil {
		log.Fatalf("Invalid input: %s", err)
	}

	// Read the content of the file
	argsArr := strings.Split(strings.ReplaceAll(fontStr, "\\n", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("../../internal/asciiart/fonts/" + banner + ".txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		defer readFile.Close()

	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted")
		return ""
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	return asciiartfs.PrintBanners(argsArr, arr)
}
