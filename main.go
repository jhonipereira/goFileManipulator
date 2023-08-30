package main

import (
	"fmt"
	"jhonidev/go/fileManip/fileutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"

	c, err := fileutils.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v \n Double the original: %v\n%v", c, c, c)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Println("ERROR: $v", err)
	}
}
