package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLineWithCallback(file string, callback func(string) error) error {
	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	lineNo := -1
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if err := callback(line); err != nil {
			return fmt.Errorf("callback error: %v, lineNo: %d", err, lineNo)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan file error: %v", err)
	}
	return nil
}
