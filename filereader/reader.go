package filereader

import (
	"bufio"
	"io/fs"
	"os"
)

func ReadLineByLine(path string) []string {
	file, err := os.OpenFile(path, os.O_RDONLY, fs.ModeAppend)
	if err != nil {
		panic("File not found")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
