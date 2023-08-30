package fileutils

import "os"

func ReadTextFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		//couldn't read the file
		return "", err
	} else {
		//read ok
		return string(content), nil
	}
}
