package metricmodules

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/SerdaOzun/Gopham/internal/storage"
)

// Credit for algorithm goes to: https://gist.github.com/shiva27/1432290

func ComputeSourceLoC(root string) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	//LoC of all Go files in the root directory
	totalLoc := 0
	for _, file := range files {
		scanFile(file, &totalLoc)
	}
	storage.SetTotalLoC(totalLoc)
	// fmt.Printf("%v total lines of code\n", totalLoc)
}

func scanFile(filepath string, totalLoc *int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	loc := 0
	insideComment := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if insideComment {
			if commentFinished(line) {
				line = strings.TrimSpace(line[strings.Index(line, "*/")+2:])
				insideComment = false
				if line == "" || strings.HasPrefix(line, "//") {
					continue
				}
			} else {
				continue
			}
		}
		if isSourceLoC(line) {
			loc++
		}
		if commentBegan(line) {
			insideComment = true
		}
	}
	*totalLoc += loc
	storage.AddFileLoCItem(filepath, loc)
	// fmt.Printf("%v : %v\n", filepath, loc)
}

func commentBegan(line string) bool {
	index := strings.Index(line, "/*")
	if index < 0 {
		return false
	}
	quoteStartIndex := strings.Index(line, "\"")
	if quoteStartIndex != -1 && quoteStartIndex < index {
		for quoteStartIndex > -1 {
			line = line[quoteStartIndex+1:]
			quoteEndIndex := strings.Index(line, "\"")
			line = line[quoteEndIndex+1:]
			quoteStartIndex = strings.Index(line, "\"")
		}
		return commentBegan(line)
	}
	return !commentFinished(line[index+2:])
}

func commentFinished(line string) bool {
	index := strings.Index(line, "*/")
	if index < 0 {
		return false
	} else {
		lineSubString := strings.TrimSpace(line[index+2:])
		if lineSubString == "" || strings.HasPrefix(lineSubString, "//") {
			return true
		}
		if commentBegan(lineSubString) {
			return false
		} else {
			return true
		}
	}
}

func isSourceLoC(line string) bool {
	isSourceLoC := false
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "//") {
		return isSourceLoC
	}
	if len(line) == 1 {
		return true
	}
	index := strings.Index(line, "/*")
	if index != 0 {
		return true
	} else {
		for len(line) > 0 {
			line = line[index+2:]
			endCommentPosition := strings.Index(line, "*/")
			if endCommentPosition < 0 {
				return false
			}
			if endCommentPosition == len(line)-2 {
				return false
			} else {
				lineSubString := strings.TrimSpace(line[endCommentPosition+2:])
				if lineSubString == "" || strings.Index(lineSubString, "//") == 0 {
					return false
				} else {
					if strings.HasPrefix(lineSubString, "/*") {
						line = lineSubString
						continue
					}
					return true
				}
			}
		}
	}
	return isSourceLoC
}
